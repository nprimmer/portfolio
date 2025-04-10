# portfolio
This repository is being put in place as a portfolio to showcase some examples of my work.

## Articles (articles)

This section holds any blog-style articles or writing that I'm able to share.
These are intentionally not dated, placed in any sort of order, or scoped
specifically to any type of technical pursuit.

### Crafting The Challenge

[Link Here](articles/crafting-the-challenge/index.md)

This article, aimed at an audience of Challenge Series administrators,
describes the process for ideation and creation of a new Challenge. The
expectation is that the readers of this article are technical contributors to
the Challenge Series, likely new volunteer administrators, or sponsor employees
who are crafting sponsored challenges.

## Challenge Series (challenge-series)

Anything under this directory is something I built as part of the GopherCon Challenge Series.

GopherCon Challenge Series is a community Capture-The-Flag event, tied to GopherCon each year since 2023. It features a variety of security, programming, and logic puzzles that are intended to test the capabilities of its contestants.

### MCP Intrusion

[Code Here](challenge-series/mcp-intrusion/)

Themed around Tron

The focus of this challenge is for the users to learn different ways of inspecting and manipulating Docker containers and images to find Flags.

### Neon Nights

[Code Here](challenge-series/neon-nights/)

Themed around Miami Vice

The focus of this challenge is for the users to inspect the contents of a virtual person's wallet, using information found about the person to try and reverse-engineer their ATM PIN, emulating basic social engineering skills.

### Nova Messages

[Code Here](challenge-series/nova-messages/)

Themed around Short Circuit

The focus of this challenge is to traverse an open API of a company's internal messaging system, perusing the messages available to identify potential security risks that will allow them to access authenticated endpoints. By astutely following the messages throughout the system, they'll discover that a specific user has poor security practices, uses simple, easily guessable passwords, and also that the IT department has implemented a system to reduce the risk from that user by resetting their password when unusual activity is detected. 

### OCP Service

[Code Here](challenge-series/ocp-service/)

Themed around Robocop

The focus of this challenge is for the player to open a websocket connection to a remote service. The remote service provides each open connection with a series of math problems of escalating size and complexity. The user needs to determine if each problem has the correct answer, and either verify its correctness, or provide the correct answer, emulating the concept of calibrating sensor data to be more accurate based on feedback. Complicating the player's efforts is that the calibration falls apart the longer the session is open, meaning that they'll lose progress if they try to perform the calibration manually, and instead need to focus on building a programmatic solution to the problem.

### WOPR

[Code Here](challenge-series/wopr/)

Themed around War Games

The focus of this challenge is for the player to open a websocket connection to a remote service. This service emulates the WOPR supercomputer from the movie War Games, and offers the ability to initiate the simulation of Global Thermonuclear War, and to play Tic-Tac-Toe.
Not available in this instance is the adjacent Wiki that was running during the competition, which provided information about the Dr Falken analogue, Dr Alex Winters. Similar to the work done in the movie, the password for access was found by researching Dr Winters in the Wiki after seeing the name mentioned upon first connecting to WOPR.