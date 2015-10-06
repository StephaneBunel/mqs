# mqs - MariaDB/MySQL Query Sniper

`mqs` is a tool for MariaDB and MySQL databases. Its mission is to kill queries that run for too long.

## mqs is flexible

`mqs` isn't a fool sniper killing blindly. He acts according to your contacts.
Because you are the best person to know if the request you wrote must be killed after a while, 
you set yourself the timeout for it.

  * As a query writer, you decide yourself if a query needs to be killed (avoid to kill INSERT nor UPDATE statements).
  * As a query writer, you decide yourself of the timeout in seconds.

## How does it works

Once `msq` is connected to the database it monitors every running queries looking for SQL statement that starts with a C-style comment. When found it acts according to given instructions.

Let see an example:

```
/* -mqs-timeout=5 */ SELECT bigdata FROM bigtable;
```
Do not add spaces around '='.

## Available instructions

  * -mqs-timeout=[integer]

This tag instruct `mqs` to kill the query if it run longer than the specified timeout in seconds.
When the query is killed, `mqs` output a trace with the SQL statement.

## Running mqs

```
MQS_DSN="root:@/mysql" bin/mqs
```

or

```
bin/mqs -dsn="root:@/mysql"
```
This last form has precedence over the first.

Note that the database user MUST have PROCESS and SUPER privileges.

## Test in mysql console

You must run mysql client with the `-c` option otherwise it will remove all comments passed to the server.

Open two terminals.
In the first, run `mqs`.
In the second, run:

```
time mysql -c -u <user> -p <password> -e "/* -mqs-timeout=5 */ select sleep(60);"
```

Without `mqs` this query sleep 60 seconds before returns. With `mqs` running, it will be killed after five seconds:

```
+-----------+
| sleep(60) |
+-----------+
|         1 |
+-----------+

real    0m6.867s
user    0m0.004s
sys     0m0.000s
```

## How to build mqs

`mqs` is written in [go](https://golang.org/)(lang) and build with [gb](http://getgb.io/).
To compile and install `mqs` just type `gb build` in the project root. The executable will be in the `bin` directory (`bin/mqs`)

