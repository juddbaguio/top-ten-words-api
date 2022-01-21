# top-ten-words-api

## "/" path (POST)

### Request Body Example (Content-Type: application/json; charset=utf-8"
```json
{
    "text": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
}
```

### Response Body Example
```json
[
    {
        "word": "the",
        "count": 6
    },
    {
        "word": "of",
        "count": 4
    },
    {
        "word": "ipsum",
        "count": 4
    },
    {
        "word": "lorem",
        "count": 4
    },
    {
        "word": "it",
        "count": 3
    },
    {
        "word": "and",
        "count": 3
    },
    {
        "word": "s",
        "count": 2
    },
    {
        "word": "typesetting",
        "count": 2
    },
    {
        "word": "with",
        "count": 2
    },
    {
        "word": "text",
        "count": 2
    }
]
```
