# CLI Todo Application

A lightweight and efficient command-line tool to manage your tasks, built with Go.

## Features

- Add, view, update, and delete tasks.
- Mark tasks as completed or pending.
- Persistent storage using a local file/database.
- User-friendly and intuitive CLI interface.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/todo-cli.git
   cd todo-cli
   ```

2. Build the application
   ```bash
   go build
   ```

## Usage

1. Adding a task

   ```bash
   ./todo -add "Write a README for Todo App"
   ```

2. Updating a task

   ```bash
   ./todo -edit "Write a README for Todo App"
   ```

3. List all task

   ```bash
   ./todo -list id:new_title
   ```

4. Deleting a task

   ```bash
   ./todo -del id
   ```

5. Toggle status
   ```bash
   ./todo -toggle id
   ```
