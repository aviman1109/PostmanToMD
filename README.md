# PostmanToMD

Convert Postman Collection to gitbook api file

This is a tool for postman collection to build a GitBook folder.
I use [gitbook plugin](https://www.npmjs.com/package/gitbook-plugin-api) to build gitbook file.
At the same time, it will build a Sammary.md file for you.

But to execute this file you need to add `GOOGLE_APPLICATION_CREDENTIALS` in your envirment variable.

```bash
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/json/your_google_translate_api.json"
```

For execute in command:

```bash
./GO_PostmanToMD your_postman_collection.json
```
