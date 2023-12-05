package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var authenticated bool
var username string
var password string
var todos []string

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple CLI todo list with authentication",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd != registerCmd && !authenticated && cmd != loginCmd && cmd != addCmd && cmd != listCmd {
			fmt.Println("Authentication required. Please log in.")
			os.Exit(1)
		}
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to access the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		authenticated = authenticate(username, password)
		if authenticated {
			fmt.Println("Login successful!")
		} else {
			fmt.Println("Invalid credentials. Login failed.")
		}
	},
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {
		registerUser(username, password)

		fmt.Println("User registered successfully!")
	},
}

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a task to the todo list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !authenticated {
			fmt.Println("Authentication required. Please log in.")
			return
		}

		task := strings.Join(args, " ")
		saveTask(task)
		fmt.Printf("Added task: %s\n", task)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		if !authenticated {
			fmt.Println("Authentication required. Please log in.")
			return
		}
		fmt.Println("Todo List:")
		displayTasks()
	},
}

func readAuthStatus() bool {
	file, err := os.Open("auth_status.txt")
	if err != nil {
		return false
	}
	defer file.Close()

	var authStatus bool
	_, err = fmt.Fscanf(file, "%t", &authStatus)
	if err != nil {
		return false
	}

	return authStatus
}

func writeAuthStatus(auth bool) error {
	file, err := os.OpenFile("auth_status.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%t", auth)
	if err != nil {
		return err
	}

	return nil
}

func authenticate(username, password string) bool {
	// Simulating authentication - using text file for user credentials
	file, err := os.Open("users.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, "|")
		if data[0] == username {
			err := bcrypt.CompareHashAndPassword([]byte(data[1]), []byte(password))
			if err == nil {
				authenticated = true
				err = writeAuthStatus(authenticated)
				if err != nil {
					fmt.Println("Error:", err)
				}
				return true
			}
		}
	}

	return false
}

func registerUser(username, password string) {
	// Simulating registration - appending user credentials to a text file
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	data := fmt.Sprintf("%s|%s\n", username, string(hashedPassword))
	file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	authenticated = true // Set authenticated to true after successful registration

	err = writeAuthStatus(authenticated)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func saveTask(task string) {
	// Simulating task addition - appending tasks to a text file
	file, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(task + "\n")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func displayTasks() {
	// Simulating task display - reading tasks from a text file
	file, err := os.Open("tasks.txt")
	if err != nil {
		fmt.Println("No tasks found.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		task := scanner.Text()
		fmt.Println("-", task)
	}
}

func main() {
	authenticated = readAuthStatus()

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)

	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Your username")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Your password")

	registerCmd.Flags().StringVarP(&username, "username", "u", "", "Your username")
	registerCmd.Flags().StringVarP(&password, "password", "p", "", "Your password")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
