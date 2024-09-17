# JsonMockServer (jms)

JsonMockServer (jms) is a Go application built with the Gin framework that automatically creates a mock server based on JSON files in your project directory. It's designed to simplify the process of setting up mock APIs for testing purposes.

## Features

- Automatically detects JSON files in the application's directory
- Creates API endpoints based on the folder structure of JSON files
- Supports both GET and POST methods
- Easy to set up and use
- Built with Go and Gin for high performance

## How It Works

1. The application scans the directory where it's run for JSON files.
2. It creates a root and endpoints based on the folder structure of these files.
3. The HTTP method is determined by the file name:
   - Files ending with `*get.json` are treated as GET endpoints
   - Files ending with `*post.json` are treated as POST endpoints
4. A mock server is created using Gin, ready for your testing needs.

## Installation

### Releases 
You can download from [Releases](https://github.com/Shahryar-Pirooz/jms/releases/tag/v0.1.0-beta) and then run:

```
./jms [:PORT]

```

### Go
To install jms, make sure you have Go installed on your system, then run:

```
git clone github.com/Shahryar-Pirooz/jms
go run . [:PORT]

```
## Usage

1. Place your JSON files in the directory where you want to run the mock server.
2. Name your files according to the desired HTTP method (e.g., `users_get.json`, `create_user_post.json`).
3. Run the jms application in that directory:

```
jms
```

4. The mock server will start, and you can now make requests to the endpoints it created.

To run the server on a specific port, you can specify the port number after the command:

```
jms :3000
```

This will start the server on port 3000.

## Endpoint Addressing

The endpoint addresses are created based on the folder structure and file names. Here are some examples:

- `/project/dir/_get.json` creates an endpoint at `/`
- `/project/dir/users_get.json` creates an endpoint at `/users/`
- `/project/dir/posts/id_get.json` creates an endpoint at `/posts/id`

## Example

If you have the following file structure:

```
/project_root
  /api
    /users
      getUsers_get.json
      createUser_post.json
    /products
      _get.json
```

jms will create the following endpoints:

- GET /api/users/getUsers
- POST /api/users/createUser
- GET /api/products

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

This is my first Golang application. Feedback and suggestions are greatly appreciated!
