package methods

import (
	"fmt"
	"strings"
)

// Method represents method (func with receiver)
type Method interface {
	GetMethodName() string
	GetReceiverDeclaration() string
	GetArgsDeclaration() string
	GetReturnValuesDeclaration() string
	GetBody() string
	GetDoc(methodName string) string
}

// receiverMethod

type receiverMethod struct {
	receiverDeclaration string
}

// GetReceiverDeclaration returns receiver declaration
func (m receiverMethod) GetReceiverDeclaration() string {
	return m.receiverDeclaration
}

func newReceiverMethod(decl string) *receiverMethod {
	return &receiverMethod{
		receiverDeclaration: decl,
	}
}

// structMethod

type structMethod struct {
	*receiverMethod
	structTypeName string
}

func newStructMethod(receiverArgName, structTypeName string) structMethod {
	receiverDecl := fmt.Sprintf("%s %s", receiverArgName, structTypeName)
	return structMethod{
		receiverMethod: newReceiverMethod(receiverDecl),
		structTypeName: structTypeName,
	}
}

// namedMethod

type namedMethod struct {
	name string
	doc  string
}

func newNamedMethod(name string) namedMethod {
	return namedMethod{
		name: name,
	}
}

// GetMethodName returns name of method
func (m namedMethod) GetMethodName() string {
	return m.name
}

// GetDoc returns default doc
func (m namedMethod) GetDoc(methodName string) string {
	if m.doc != "" {
		return m.doc
	}

	return fmt.Sprintf(`// %s is an autogenerated method
	// nolint: dupl`, methodName)
}

func (m *namedMethod) setDoc(doc string) {
	m.doc = doc
}

// oneArgMethod

type oneArgMethod struct {
	argName     string
	argTypeName string
}

func (m oneArgMethod) getArgName() string {
	return m.argName
}

// GetArgsDeclaration returns declaration of arguments list for func decl
func (m oneArgMethod) GetArgsDeclaration() string {
	return fmt.Sprintf("%s %s", m.getArgName(), m.argTypeName)
}

func newOneArgMethod(argName, argTypeName string) oneArgMethod {
	return oneArgMethod{
		argName:     argName,
		argTypeName: argTypeName,
	}
}

type nArgsMethod struct {
	args []oneArgMethod
}

func (m nArgsMethod) getArgName(n int) string {
	return m.args[n].getArgName()
}

// GetArgsDeclaration returns declaration of arguments list for func decl
func (m nArgsMethod) GetArgsDeclaration() string {
	decls := []string{}
	for _, a := range m.args {
		decls = append(decls, a.GetArgsDeclaration())
	}
	return strings.Join(decls, ",")
}

func newNArgsMethod(args ...oneArgMethod) nArgsMethod {
	return nArgsMethod{
		args: args,
	}
}

// noArgsMethod

type noArgsMethod struct{}

// GetArgsDeclaration returns declaration of arguments list for func decl
func (m noArgsMethod) GetArgsDeclaration() string {
	return ""
}

// errorRetMethod

type errorRetMethod struct{}

func (m errorRetMethod) GetReturnValuesDeclaration() string {
	return "error"
}

// constBodyMethod

type constBodyMethod struct {
	body string
}

// GetBody returns const body
func (m constBodyMethod) GetBody() string {
	return m.body
}

func newConstBodyMethod(format string, args ...interface{}) constBodyMethod {
	return constBodyMethod{
		body: fmt.Sprintf(format, args...),
	}
}

// constRetMethod

type constRetMethod struct {
	ret string
}

func (m constRetMethod) GetReturnValuesDeclaration() string {
	return m.ret
}

func newConstRetMethod(ret string) constRetMethod {
	return constRetMethod{
		ret: ret,
	}
}
