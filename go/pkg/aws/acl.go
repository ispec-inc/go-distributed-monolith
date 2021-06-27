package aws

const (
	ACLPrivate           = acl("private")
	ACLPublicRead        = acl("public-read")
	ACLPublicReadWrite   = acl("public-read-write")
	ACLExecRead          = acl("aws-exec-read")
	ACLAuthenticatedRead = acl("authenticated-read")
)

type acl string
