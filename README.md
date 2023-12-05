# Simple CLI Todo List with Authentication

This CLI application allows you to manage your todo list with basic authentication features. It uses command-line inputs to register users, log in, add tasks, and list all tasks.

## Installation

To use this application, ensure you have Go installed. Clone the repository and navigate to the project directory:

```bash 
    git clone https://github.com/MamangRust/todo_list_auth_go.git todo_cli_app
    cd todo_cli_app
```

### Then, build the application:

```bash 
    go build -o todo
```

## Usage

### Register a New User

To register a new user:

```bash 
    ./todo register -u username -p password
```

Replace `username` and `password` with your desired credentials.

### Log In

To log in as a registered user:

```bash
    ./todo login -u username -p password
```

### Add Tasks

After logging in, add tasks to your todo list:

```bash 
./todo add "Your task description here"
```



### List Tasks

List all tasks in your todo list:

```bash 
./todo list

```


## Note

- Ensure that you have write permissions in the directory to create `users.txt` and `tasks.txt` files for user authentication and task storage.
- The authentication uses basic text file storage for demonstration purposes. In a production environment, use a secure and reliable authentication method.
