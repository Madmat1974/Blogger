# Gator

A multi-player command line tool for aggregating RSS feeds and viewing the posts.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `gator` with:

go install github.com/Madmat1974/gator@latest

Config
Create a .gatorconfig.json file in your home directory with the following structure:

{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}

Replace the values with your database connection string.

Usage
Create a new user:

gator register <name>

Add a feed:

gator addfeed <url>

Start the aggregator:

gator agg 30s

View the posts:

gator browse [limit]

There are a few other commands you'll need as well:

gator login <name> - Log in as a user that already exists
gator users - List all users
gator feeds - List all feeds
gator follow <url> - Follow a feed that already exists in the database
gator unfollow <url> - Unfollow a feed that already exists in the database