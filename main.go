/*
	Author: 	 Dylen Greenenwald
	UIN:		 658146889
	Date:        2023-11-11
	Description: A simple password manager that relies heavily on maps and
				 slices. A user of this manager can load it with a properly
				 formatted file (in this case, hard coded to be passwordVault).
				 The user can then list the contents of the password manager,
				 add an entry, remove an entry, or terminate the process.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Entry represents an entry in the password manager.
type Entry struct {
	User     string  // The user who is seeking to store credentials for site.
	Password string  // The password for the user at the site.
}

type EntrySlice []Entry

// The data structure that represents the password manager.
var passwordMap map[string]EntrySlice

// initialize before main()
func init() {
	passwordMap = make(map[string]EntrySlice)
}

// findEntrySlice finds the entry slice for a given site.
func findEntrySlice(site string) (EntrySlice, bool) {
	// Accessing a map in Go returns an entry if it is found, and a zero
	// value if it is not found. Additionally, it will return a flag
	// indicating whether or not the entry was found.
	entrySlice, found := passwordMap[site]
	return entrySlice, found
}

// setEntrySlice sets the entry slice for a site.
func setEntrySlice(site string, entrySlice EntrySlice) {
	passwordMap[site] = entrySlice
}

// find finds a user within an entry slice.
func find(user string, entrySlice EntrySlice) (int, bool) {
	// For each entry in the entry slice, check if the user matches the
	// user we are looking for. If it does, return the index of the entry
	// and a flag indicating that the user was found.
	for i, entry := range entrySlice {
		if entry.User == user {
			return i, true
		}
	}

	// If control reaches here, the user was not found.
	return -1, false
}

// pmList lists the contents of the password manager in a tabular format.
func pmList() {
	fmt.Println("Site\tUser\tPassword")
	for site, entrySlice := range passwordMap {
		for _, entry := range entrySlice {
			fmt.Printf("%s\t%s\t%s\n", site, entry.User, entry.Password)
		}
	}
}

// pmAdd adds an entry for a site and user, ensuring there are no duplicate
// entries.
func pmAdd(site, user, password string) {
	// Populate entrySlice with existing entries for the site if it exists;
	// update a corresponding boolean.
	entrySlice, found := findEntrySlice(site)

	// If there is not an entry with this site, create a new entrySlice to
	// hold this entry and future entries with this site.
	if !found {
		entrySlice = make(EntrySlice, 0)
	}

	// Ignore the index of where, but check if the user was found. If so, 
	// prevent adding a duplicate entry.
	_, userFound := find(user, entrySlice)
	if userFound {
		fmt.Println("add: duplicate entry")
		return
	}

	// Since the entry doesn't exist, we can create it and append it to the
	// entrySlice.
	entry := Entry{
		Site:     site,
		User:     user,
		Password: password,
	}
	entrySlice = append(entrySlice, entry)
	setEntrySlice(site, entrySlice)
}

// pmRemove removes an entry for a specific site and user.
func pmRemove(site, user string) {
	// Check if there is an entry for the site
	entrySlice, found := findEntrySlice(site)
	if !found {
		fmt.Println("remove: user/site not found")
		return
	}

	// Check if there is a user inside of the existing slice
	index, userFound := find(user, entrySlice)
	if !userFound {
		fmt.Println("remove: user not found")
		return
	}

	// Remove the entry from the slice and update the entry slice for the
	// site.
	entrySlice = append(entrySlice[:index], entrySlice[index+1:]...)
	setEntrySlice(site, entrySlice)
}

// pmRemoveSite removes the entire site if there is only a single user at
// that site.
func pmRemoveSite(site string) {
	entrySlice, found := findEntrySlice(site)
	
	// If there are no entries for this site, there is nothing to remove.
	if !found {
		fmt.Println("remove: site not found")
		return
	}

	// If there is a single entry for this site, remove the site. Else,
	// print an error.
	if len(entrySlice) == 1 {
		delete(passwordMap, site)
	} else {
		fmt.Println("remove: attempted to remove multiple errors")
	}
}

// pmRead reads the passwordVault file to initialize the password manager.
func pmRead() {
	// Define streams for reading the passwordVault file.
	file, err := os.Open("passwordVault")

	// If we encounter an error, stop reading the file.
	if err != nil {
		fmt.Println("Error reading passwordVault:", err)
		return
	}
	// Queue up closing the file for when we are done reading it.
	defer file.Close()

	// Read the file line by line, splitting each line into its parts.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		// If the line is improperly formatted, just skip it.
		if len(parts) == 3 {
			pmAdd(parts[0], parts[1], parts[2])
		}
	}
}

// pmWrite writes the passwordVault file.
func pmWrite() {
	file, err := os.Create("passwordVault")
	if err != nil {
		fmt.Println("Error writing passwordVault:", err)
		return
	}
	defer file.Close()

	for site, entrySlice := range passwordMap {
		for _, entry := range entrySlice {
			_, err := file.WriteString(fmt.Sprintf("%s %s %s\n", site, entry.User, entry.Password))
			if err != nil {
				fmt.Println("Error writing to passwordVault:", err)
				return
			}
		}
	}
}

// loop reads commands and processes them.
func loop() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a command (l/a/r/x): ")
		scanner.Scan()
		command := scanner.Text()
		switch command {
		case "l":
			pmList()
		case "a":
			fmt.Print("Enter site, user, password (separated by spaces): ")
			scanner.Scan()
			input := scanner.Text()
			parts := strings.Split(input, " ")
			if len(parts) == 3 {
				pmAdd(parts[0], parts[1], parts[2])
			}
		case "r":
			fmt.Print("Enter site and user to remove (separated by a space): ")
			scanner.Scan()
			input := scanner.Text()
			parts := strings.Split(input, " ")
			if len(parts) == 2 {
				pmRemove(parts[0], parts[1])
			}
		case "x":
			pmWrite()
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Please enter l/a/r/x.")
		}
	}
}

func main() {
	pmRead()
	loop()
}
