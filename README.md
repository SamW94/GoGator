# GoGator (GoGator)

Gator is an RSS feed aggregator build in Go. It is a CLI tool that allows users to: 

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a local PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of aggregated posts in the terminal, with a link to the full post

The intention of this toy project is to demonstrate an understanding of: 

- Integrating a Golang application with a PostgreSQL database.
- SQL queries and migrations. 
- Writing a long-running service that continuously makes HTTP requests and processes responses.
- Writing a long-running service that continuously retrieves and inserts data to a database.
- Using the [sqlc](https://sqlc.dev/) and [goose](https://github.com/pressly/goose) tools for creating typesafe SQL in Go.

## Getting Started

