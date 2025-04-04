### Review config managment

**Issue**: Hardcoding the configuration means that any changes require recompilation of the application. This can lead to inconvenience and error-prone scenarios, especially in production.

**Suggestion**: Use a configuration management library like [Viper](https://github.com/spf13/viper) or simply use environment variables with Go's built-in `os.Getenv`. These libraries allow for loading configurations from various sources (e.g., JSON, YAML, ENV variables) and make managing different environments easier (development, testing, production).

**Reason**: Increases flexibility and allows for changing configurations without modifying the source code or recompiling the application.

### Review custom client

It's great to have a custom client for testing the behavior, but it would be nice to have unit tests as well. This would provide more informative feedback, especially since tools like Postman and Insomnia handle gRPC requests effectively. I appreciate the approach you've taken.

### Review proto

#### 1. Package Naming Convention

**Issue**: The package name does not include a version suffix, which can lead to confusion and potential breaking changes in the future.

**Suggestion**: Consider using a versioning scheme for the package name, such as `ad.v1`.

**Reason**: Versioned package names help maintain backward compatibility and provide a clear indication of changes across different iterations of your API. This is especially crucial for APIs that may evolve over time.

#### 2. Message Reusability

**Issue**: The `AdRequest` message is used in multiple places, but it's only defined once, which could lead to inconsistencies if modifications are needed.

**Suggestion**: Rename the existing `AdRequest` or create more specific request messages. For example, you could have `ReadAdRequest` and `ServeAdRequest`.

**Reason**: Having distinct message types makes it clear what each message is intended to do and helps avoid confusion regarding input requirements. This practice enhances readability and reduces the risk of errors in the future.

### 3. Input Validation

**Issue**: Currently, validation logic is implemented at the service level in Go. While this approach is functional, it lacks portability, which is one of the key advantages of using Protocol Buffers.

**Suggestion**: Implement [ProtoValidate](https://github.com/mengzhuo/proto-validate) for automatic validation based on your message definitions.

**Reason**: Having validation at the message level helps to reduce the burden on the service layer and promotes reusability of validation logic. This encapsulates constraints directly within the schema, ensuring consistency across different services and clients.

**Documentation**: Adding comments to each message and service method can significantly improve the understandability of the API, especially for new developers or external users.

The protobuf code provided is a solid foundation for your gRPC service but can greatly benefit from some structural refinements and best practices. Implementing versioned package naming, enhancing message clarity and validation, and adopting clear documentation practices will help improve maintainability, reliability, and ease of use.

### Review global

The application is functioning well, but it would be beneficial to include a Docker Compose setup to easily launch containers and services.

The lack of tests makes it challenging to verify each case locally; having some unit or integration tests would significantly improve the situation.

I'm puzzled by the inconsistency in error handling; for example, in `ConnectToMongoDB()`, you handle errors appropriately, but in `DisconnectFromMongoDB()`, you decide to skip error handling by using `fatal`. This approach is generally not recommended.

Additionally, the project architecture seems a bit unusual. While you are using a layered pattern and have well-separated database functions within the `db` package for connecting and disconnecting, it appears that you are calling MongoDB directly from the service. It would be better to encapsulate these operations within the `db` package for better abstraction and maintainability.
