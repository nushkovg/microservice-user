package design

// Use . imports to enable the DSL
import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Define default description and default global property values
var _ = API("user", func() {
	Title("The user microservice")
	Description("A service that provides basic access to the user data")
	Version("1.0")
	Scheme("http")
	Host("localhost:8080")
})

// Resources group related API endpoints together.
var _ = Resource("user", func() {
	BasePath("/users")
	DefaultMedia(UserMedia)
	// Do not setup security here!

	// Actions define a single API endpoint
	Action("create", func() {
		Description("Creates user")
		Routing(POST(""))
		Payload(UserPayload)
		Response(Created, UserMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("get", func() {
		Description("Get user by id")
		Routing(GET("/:userId"))
		Params(func() {
			Param("userId", String, "User ID")
		})
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("getMe", func() {
		Description("Retrieves the user information for the authenticated user")
		Routing(GET("/me"))
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("update", func() {
		Description("Update user")
		Routing(PUT("/:userId"))
		Params(func() {
			Param("userId", String, "User ID")
		})
		Payload(UserPayload)
		Response(OK, UserMedia)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("find", func() {
		Description("Find a user by email+password")
		Routing(POST("find"))
		Payload(CredentialsPayload)
		Response(OK, UserMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("findByEmail", func() {
		Description("Find a user by email")
		Routing(POST("find/email"))
		Payload(EmailPayload)
		Response(OK, UserMedia)
		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("verify", func() {
		Description("Verify a user by token")
		Routing(GET("verify"))
		Params(func() {
			Param("token", String, "Token")
		})
		Response("OK", func() {
			Description("User is verified")
			Status(200)
			Media("plain/text")
		})
		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
	Action("resetVerificationToken", func() {
		Description("Reset verification token")
		Routing(POST("verification/reset"))
		Payload(EmailPayload)
		Response("OK", func() {
			Description("Verification token reset")
			Status(200)
			Media(ResetTokenMedia)
		})
		Response(BadRequest, ErrorMedia)
		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})

// UserMedia defines the media type used to render user.
var UserMedia = MediaType("application/vnd.goa.user+json", func() {
	TypeName("users")
	Reference(UserPayload)

	Attributes(func() {
		Attribute("id", String, "Unique user ID")
		Attribute("email")
		Attribute("roles")
		Attribute("externalId")
		Attribute("active")
		Attribute("organizations")
		Attribute("namespaces")
		Required("id", "email", "roles", "externalId", "active")
	})

	View("default", func() {
		Attribute("id")
		Attribute("email")
		Attribute("roles")
		Attribute("externalId")
		Attribute("active")
		Attribute("organizations")
		Attribute("namespaces")
	})
})

// ResetTokenMedia is returned after successful reset of the verification token
var ResetTokenMedia = MediaType("ResetTokenMedia", func() {
	TypeName("ResetToken")
	Attributes(func() {
		Attribute("id", String, "User ID")
		Attribute("email", String, "User email")
		Attribute("token", String, "New token")
		Required("id", "email", "token")
	})
	View("default", func() {
		Attribute("id")
		Attribute("email")
		Attribute("token")
	})
})

// UserPayload defines the payload for the user.
var UserPayload = Type("UserPayload", func() {
	Description("UserPayload")

	Attribute("email", String, "Email of user", func() {
		Format("email")
	})
	Attribute("password", String, "Password of user", func() {
		MinLength(6)
		MaxLength(30)
	})
	Attribute("roles", ArrayOf(String), "Roles of user")
	Attribute("organizations", ArrayOf(String), "List of organizations to which this user belongs to")
	Attribute("namespaces", ArrayOf(String), "List of namespaces this user belongs to")
	Attribute("externalId", String, "External id of user")
	Attribute("active", Boolean, "Status of user account", func() {
		Default(false)
	})
	Attribute("token", String, "Token for email verification")

	Required("email")
})

// CredentialsPayload defines the payload for the credentials.
var CredentialsPayload = Type("Credentials", func() {
	Description("Email and password credentials")
	Attribute("email", String, "Email of user", func() {
		Format("email")
	})
	Attribute("password", String, "Password of user", func() {
		MinLength(6)
		MaxLength(30)
	})
	Required("email", "password")
})

// EmailPayload defines the payload for the email.
var EmailPayload = Type("EmailPayload", func() {
	Description("Email payload")
	Attribute("email", String, "Email of user", func() {
		Format("email")
	})
	Required("email")
})

// Swagger UI
var _ = Resource("swagger", func() {
	Description("The API swagger specification")

	Files("swagger.json", "swagger/swagger.json")
	Files("swagger-ui/*filepath", "swagger-ui/dist")
})
