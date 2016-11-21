package main


import (
	"os"
	"fmt"
	"regexp"
	"bytes"
	"strings"
	j "encoding/json"

	"github.com/jhaynie/go-github-protobuf/github"
	_ "github.com/golang/protobuf/proto"
	d "google.golang.org/genproto/protobuf"
	_ "github.com/golang/protobuf/descriptor"
)

var re = regexp.MustCompile("[-_]([a-z])")
var re2 = regexp.MustCompile("\\d+:\"(\\w+)\"")

func stringify(v interface{}) string {
	buf, err := j.MarshalIndent(v, "", "\t")
	if err != nil {
		return ""
	}
	return string(buf)
}

func makeJSName(n string) string {
	return strings.Replace(n, ".proto", "_pb.js", -1)
}

func makeJSFilename(ver string, n string) string {
	return fmt.Sprintf("build/%s/javascript/%s", ver, makeJSName(n))
}

func getFieldName(s string) string {
	return re2.FindStringSubmatch(s)[1]
}

func repl(s string) string {
	return strings.Title(s[1:len(s)])
}

func makeCamelCase(prefix string, s string) string {
	return fmt.Sprintf("%s%s", prefix, strings.Title(re.ReplaceAllStringFunc(s, repl)))
}

func appendFile(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(text)
	if err != nil {
		panic(err)
	}
}

func getNestedType(name string, desc []*d.DescriptorProto) *d.DescriptorProto {
	for _, p := range desc {
		if *p.Name == name {
			return p
		}
	}
	return nil
}

func dedot(name string) string {
	tok := strings.Split(name, ".")
	if len(tok) > 1 {
		return tok[len(tok)-1]
	}
	return name
}

func main() {

	version := os.Args[1]

	processedfiles := make(map[string]bool)

	descriptors, filenames := github.GetGithubDescriptors()

	for k, fielddef := range descriptors {
		var buffer bytes.Buffer
		dfn := makeJSFilename(version, filenames[k])

		if processedfiles[dfn] == false {
			buffer.WriteString("\n\n// patched by github-protobuf to add toJSON and fromJSON methods\n\n")
			buffer.WriteString("function _toBool (obj) {\n")
			buffer.WriteString("	if (typeof(obj) === 'boolean') { return obj; }\n")
			buffer.WriteString("	if (typeof(obj) === 'string') { return obj === 'true'; }\n")
			buffer.WriteString("	if (typeof(obj) === 'number') { return obj > 0; }\n")
			buffer.WriteString("	return false;\n")
			buffer.WriteString("};\n\n")
			processedfiles[dfn] = true
		}

		buffer.WriteString(fmt.Sprintf("\n\n// %s\n", k))
		buffer.WriteString(fmt.Sprintf("proto%s.prototype.fromJSON = function(obj) {\n", k))
		// fmt.Println(k, "=>", *fielddef.Name, "=>", filenames[k], dfn)
		for _, field := range fielddef.Field {
			fieldname := *field.Name
			fo := field.GetOptions()
			if fo != nil {
				fieldname = getFieldName(fo.String())
			}
			setter := makeCamelCase("set", fieldname)
			switch *field.Type {
				case d.FieldDescriptorProto_TYPE_MESSAGE: {
					md := descriptors[field.GetTypeName()]
					if md == nil {
						if strings.HasSuffix(field.GetTypeName(), "Entry") {
							// this is a map
							mapGetter := makeCamelCase("get", fmt.Sprintf("%sMap", field.GetName()))
							nfn := dedot(field.GetTypeName())
							nt := getNestedType(nfn, fielddef.GetNestedType())
							// kt := nt.GetField()[0]
							vt := nt.GetField()[1]
							buffer.WriteString(fmt.Sprintf("	if ('%s' in obj) {\n", fieldname))
							buffer.WriteString(fmt.Sprintf("		var m = this.%s();\n", mapGetter))
							buffer.WriteString(fmt.Sprintf("		Object.keys(obj.%s).forEach(function(k) {\n", fieldname))
							if (*vt.Type == d.FieldDescriptorProto_TYPE_MESSAGE) {
								md2 := descriptors[vt.GetTypeName()]
								nfn2 := makeJSName(filenames[vt.GetTypeName()])
								buffer.WriteString(fmt.Sprintf("			var %s = require('./%s').%s;\n", *md2.Name, nfn2, *md2.Name))
								buffer.WriteString(fmt.Sprintf("			var %sInstance = new %s();\n", *md2.Name, *md2.Name))
								buffer.WriteString(fmt.Sprintf("			var v = %sInstance.fromJSON(obj.%s[k]));\n", *md2.Name, fieldname))
								buffer.WriteString("			m.set(k, v);\n")
							} else {
								buffer.WriteString(fmt.Sprintf("			m.set(k, obj.%s[k]);\n", fieldname))
							}
							buffer.WriteString("		});\n")
							buffer.WriteString("	}\n")
						} else {
							// this is a list
							if field.GetLabel() == d.FieldDescriptorProto_LABEL_REPEATED {
								// fmt.Println(field)
								buffer.WriteString(fmt.Sprintf("	if ('%s' in obj) {\n", fieldname))
								buffer.WriteString(fmt.Sprintf("		this.%sList(obj.%s);\n", makeCamelCase("set", fieldname), fieldname))
								buffer.WriteString("	}\n")
							} else {
								fmt.Println(field.GetTypeName())
								panic("not sure what to do with this type")
							}
						}
					} else {
						nfn := makeJSName(filenames[field.GetTypeName()])
						buffer.WriteString(fmt.Sprintf("	if ('%s' in obj) {\n", fieldname))
						buffer.WriteString(fmt.Sprintf("		var %s = require('./%s').%s;\n", *md.Name, nfn, *md.Name))
						buffer.WriteString(fmt.Sprintf("		var %sInstance = new %s();\n", *md.Name, *md.Name))
						buffer.WriteString(fmt.Sprintf("		this.%s(%sInstance.fromJSON(obj.%s));\n", setter, *md.Name, fieldname))
						buffer.WriteString("	}\n")
					}
				}
				case d.FieldDescriptorProto_TYPE_STRING: {
					if field.GetLabel() == d.FieldDescriptorProto_LABEL_REPEATED {
						buffer.WriteString(fmt.Sprintf("	'%s' in obj && this.%sList(obj.%s);\n", fieldname, setter, fieldname))
					} else {
						buffer.WriteString(fmt.Sprintf("	'%s' in obj && this.%s(obj.%s);\n", fieldname, setter, fieldname))
					}
				}
				case d.FieldDescriptorProto_TYPE_BOOL: {
					if field.GetLabel() == d.FieldDescriptorProto_LABEL_REPEATED {
						buffer.WriteString(fmt.Sprintf("	'%s' in obj && this.%sList(obj.%s.map(function(i){return _toBool(i);}));\n", fieldname, setter, fieldname))
					} else {
						buffer.WriteString(fmt.Sprintf("	'%s' in obj && this.%s(_toBool(obj.%s));\n", fieldname, setter, fieldname))
					}
				}
				case d.FieldDescriptorProto_TYPE_INT32: {
					if field.GetLabel() == d.FieldDescriptorProto_LABEL_REPEATED {
						buffer.WriteString(fmt.Sprintf("	'%s' in obj && this.%sList(obj.map(function(i){return +i;}));\n", fieldname, setter, fieldname))
					} else {
						buffer.WriteString(fmt.Sprintf("	'%s' in obj && this.%s(+obj.%s);\n", fieldname, setter, fieldname))
					}
				}
				case d.FieldDescriptorProto_TYPE_BYTES: {
					if field.GetLabel() == d.FieldDescriptorProto_LABEL_REPEATED {
						buffer.WriteString(fmt.Sprintf("	'%s' in obj && this.%sList(obj.%s);\n", fieldname, setter, fieldname))
					} else {
						buffer.WriteString(fmt.Sprintf("	'%s' in obj && this.%s(obj.%s);\n", fieldname, setter, fieldname))
					}
				}
			}
		}
		buffer.WriteString("	return this;\n")
		buffer.WriteString("};\n")
		buffer.WriteString("\n")
		buffer.WriteString(fmt.Sprintf("proto%s.prototype.toJSON = function() {\n", k))
		buffer.WriteString("	var obj = this.toObject();\n")
		for _, field := range fielddef.Field {
			fieldname := field.GetJsonName()
			jsonname := field.GetName()
			// fmt.Println(field)
			switch *field.Type {
				case d.FieldDescriptorProto_TYPE_MESSAGE: {
					md := descriptors[field.GetTypeName()]
					if md == nil {
						if strings.HasSuffix(field.GetTypeName(), "Entry") {
							// this is a map
							mapGetter := makeCamelCase("get", fmt.Sprintf("%sMap", fieldname))
							// fmt.Println(stringify(field))
							buffer.WriteString(fmt.Sprintf("	if ('%sMap' in obj) {\n", fieldname))
							buffer.WriteString(fmt.Sprintf("		var %s = this.%s();\n", jsonname, mapGetter))
							buffer.WriteString(fmt.Sprintf("		obj.%s = {};\n", jsonname))
							buffer.WriteString(fmt.Sprintf("		delete obj.%sMap;\n", fieldname))
							buffer.WriteString(fmt.Sprintf("		%s.forEach(function(v, k) {\n", jsonname))
							buffer.WriteString(fmt.Sprintf("			obj.%s[k] = v;\n", jsonname))
							buffer.WriteString("		});\n")
							buffer.WriteString("	}\n")
						} else {
							if field.GetLabel() == d.FieldDescriptorProto_LABEL_REPEATED {
								buffer.WriteString(fmt.Sprintf("	if ('%sList' in obj) {\n", fieldname))
								buffer.WriteString(fmt.Sprintf("		obj.%s = obj.%sList;\n", jsonname, fieldname))
								buffer.WriteString(fmt.Sprintf("		delete obj.%s;\n", fieldname))
								buffer.WriteString("	}\n")
							} else {
								fmt.Println(field.GetTypeName())
								panic("not sure what to do with this type")
							}
						}
					} else {
						buffer.WriteString(fmt.Sprintf("	if ('%s' in obj) {\n", fieldname))
						getter := makeCamelCase("get", fieldname)
						buffer.WriteString(fmt.Sprintf("		obj.%s = this.%s().toJSON();\n", jsonname, getter))
						buffer.WriteString("	}\n")
					}
				}
				default: {
					if field.GetLabel() == d.FieldDescriptorProto_LABEL_REPEATED {
						buffer.WriteString(fmt.Sprintf("	if ('%sList' in obj) {\n", fieldname))
						buffer.WriteString(fmt.Sprintf("		obj.%s = obj.%sList;\n", jsonname, fieldname))
						buffer.WriteString(fmt.Sprintf("		delete obj.%sList;\n", fieldname))
						buffer.WriteString("	}\n")
					} else {
						buffer.WriteString(fmt.Sprintf("	if ('%s' in obj) {\n", fieldname))
						buffer.WriteString(fmt.Sprintf("		obj.%s = obj.%s;\n", jsonname, fieldname))
						buffer.WriteString(fmt.Sprintf("		delete obj.%s;\n", fieldname))
						buffer.WriteString("	}\n")
					}
				}
			}
		}
		buffer.WriteString("	return obj;\n")
		buffer.WriteString("};\n")
		buffer.WriteString("\n")
		appendFile(dfn, buffer.String())
	}

}
