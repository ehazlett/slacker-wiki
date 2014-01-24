# Wiki
Searches a Confluence wiki

# Usage
Listens on port 8080.

```
./wiki -url <your-custom-confluence-url> -auth <base64-encoded-pass>
```

For the auth, you must encode "username:password" using Base64.  For example:

```
echo "admin:secretpass" | base64
```

Then add a custom command in Slack pointing to where your instance is running.
It will use the Confluence search for results.
