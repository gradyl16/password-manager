<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT DESCRIPTION -->
<br />
<div align="center">
<h3 align="center">Password Manager</h3>

  <p align="center">
    A simple command line password manager written in Go.
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This project was developed as part of the curriculum for Programming Language Design & Implementation during the fall 2023 semester. Please note that this project was built for learning purposes and is not intended to be used with live credentials.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these steps.

### Prerequisites

Read [Go's documentation](https://go.dev/doc/install) on how to install it for your operating system.

### Installation

Clone the repo:

```sh
git clone https://github.com/gradyl16/password-manager.git
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Run the program:

```sh
go run main.go
```

List all entries:

```
Enter a command (l/a/r/x): l
Site    User    Password
google.com      bob     password
blah.com        me      password
example.com     dylen   123456
```

Add an entry:

```
Enter a command (l/a/r/x): a
Enter site, user, password (separated by spaces): example.com johndoe v3rySafe
```

Remove an entry:

```
Enter a command (l/a/r/x): r        
Enter site and user to remove (separated by a space): example.com johndoe
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Dylen Greenenwald - dgree21@uic.edu

Project Link: [https://github.com/gradyl16/password-manager](https://github.com/gradyl16/password-manager)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/gradyl16/password-manager.svg?style=for-the-badge
[contributors-url]: https://github.com/gradyl16/password-manager/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/gradyl16/password-manager.svg?style=for-the-badge
[forks-url]: https://github.com/gradyl16/password-manager/network/members
[stars-shield]: https://img.shields.io/github/stars/gradyl16/password-manager.svg?style=for-the-badge
[stars-url]: https://github.com/gradyl16/password-manager/stargazers
[issues-shield]: https://img.shields.io/github/issues/gradyl16/password-manager.svg?style=for-the-badge
[issues-url]: https://github.com/gradyl16/password-manager/issues
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/gradyl16
