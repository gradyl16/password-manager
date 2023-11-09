// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Entry struct {
// 	User     string
// 	Password string
// }

// type PasswordManager map[string][]Entry

// func (pm PasswordManager) pmRead() {
// 	// Read data from the "passwordVault" file and populate the passwordMap.
// 	file, err := os.Open("passwordVault")
// 	if err != nil {
// 		fmt.Println("Error reading passwordVault:", err)
// 		return
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		fields := strings.Fields(line)
// 		if len(fields) < 3 {
// 			continue
// 		}
// 		site, user, password := fields[0], fields[1], fields[2]
// 		entry := Entry{User: user, Password: password}
// 		pm[site] = append(pm[site], entry)
// 	}
// }

// func (pm PasswordManager) pmWrite() {
// 	// Write the current state of the passwordMap back to the "passwordVault" file.
// 	file, err := os.Create("passwordVault")
// 	if err != nil {
// 		fmt.Println("Error writing passwordVault:", err)
// 		return
// 	}
// 	defer file.Close()

// 	for site, entries := range pm {
// 		for _, entry := range entries {
// 			fmt.Fprintf(file, "%s %s %s\n", site, entry.User, entry.Password)
// 		}
// 	}
// }

// func main() {
// 	pm := make(PasswordManager)
// 	pm.pmRead()
// 	defer pm.pmWrite()

// 	// Implement your command loop (l, a, r, x) and associated actions here.

// 	// Example: List all entries
// 	for site, entries := range pm {
// 		fmt.Println(site)
// 		for _, entry := range entries {
// 			fmt.Printf("  User: %s, Password: %s\n", entry.User, entry.Password)
// 		}
// 	}
// }

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Global variables are allowed (and encouraged) for this project.
type EntrySlice struct {
	User     string
	Password string
}

type EntrySlices []EntrySlice

type PasswordManager map[string]EntrySlices

var passwordMap PasswordManager


// Create an Iter method to iterate over EntrySlices
func (entries EntrySlices) Iter() <-chan EntrySlice {
	ch := make(chan EntrySlice)
	go func() {
		defer close(ch)
		for _, entry := range entries {
			ch <- entry
		}
	}()
	return ch
}

// _______________________________________________________________________

// _______________________________________________________________________
// initialize before main()
// _______________________________________________________________________
func init() {
	// Initialize the passwordMap by reading data from a file like "passwordVault"
	// if err := pmRead("passwordVault"); err != nil {
	//     fmt.Printf("Error initializing password manager: %v\n", err)
	//     os.Exit(1)
	// }
}

// _______________________________________________________________________
// find the matching entry slice
// _______________________________________________________________________
func findEntrySlice(site string) (EntrySlice, bool) {
	return EntrySlice{}, false
}

// _______________________________________________________________________
// set the entrySlice for site
// _______________________________________________________________________
func setEntrySlice(site string, entrySlice EntrySlice) {
}

// _______________________________________________________________________
// find
// _______________________________________________________________________
func find(user string, entrySlice EntrySlice) (int, bool) {
	return -1, false
}

// _______________________________________________________________________
// print the list in columns
// _______________________________________________________________________
func pmList() {
	// Define the fixed column widths
	siteWidth := 20
	userWidth := 20
	passwordWidth := 20

	// Print a header for the table
	fmt.Printf("%-*s %-*s %-*s\n", siteWidth, "Site", userWidth, "User", passwordWidth, "Password")
	fmt.Println(strings.Repeat("-", siteWidth+userWidth+passwordWidth))

	// Iterate through the passwordMap and print each site, and all slices associated
	// with that site
	for site, entries := range passwordMap {
		fmt.Printf("%-*s ", siteWidth, site)
		for entry := range entries.Iter() {
			fmt.Printf("%-*s %-*s\n", userWidth, entry.User, passwordWidth, entry.Password)
		}
	}
}

// _______________________________________________________________________
//
//	add an entry if the site, user is not already found
//
// _______________________________________________________________________
func pmAdd(site, user, password string) {
}

// _______________________________________________________________________
// remove by site and user
// _______________________________________________________________________
func pmRemove(site, user string) {
}

// _______________________________________________________________________
// remove the whole site if there is a single user at that site
// _______________________________________________________________________
func pmRemoveSite(site string) {

}

// _______________________________________________________________________
// read the passwordVault
// _______________________________________________________________________
func pmRead() {
	file, err := os.Open("passwordVault")
	if err != nil {
		log.Fatal(err)
	}

	for {
		var site, user, password string
		_, err := fmt.Fscanf(file, "%s %s %s\n", &site, &user, &password)
		if err != nil {
			break
		}
		fmt.Println(site, user, password)
	}

	defer file.Close()
}

// _______________________________________________________________________
// write the passwordVault
// _______________________________________________________________________
func (pm PasswordManager) pmWrite() {
	// Write the current state of the passwordMap back to the "passwordVault" file.
	file, err := os.Create("passwordVault")
	if err != nil {
		fmt.Println("Error writing passwordVault:", err)
		return
	}
	defer file.Close()

	for site, entries := range pm {
		for _, entry := range entries {
			fmt.Fprintf(file, "%s %s %s\n", site, entry.User, entry.Password)
		}
	}
}

// _______________________________________________________________________
// do forever loop reading the following commands
//
//	  l
//	  a s u p
//	  r s
//	  r s u
//	  x
//	where l,a,r,x are list, add, remove, and exit
//	and s,u,p are site, user, and password
//
// _______________________________________________________________________
func loop() {
}

// _______________________________________________________________________
//
//	let her rip
//
// _______________________________________________________________________
func main() {
	pmRead()
	pmList()
	os.Exit(0)
}
