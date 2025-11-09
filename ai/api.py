from model import Document
from utils import load_model
from fastapi import FastAPI

app = FastAPI()

@app.get("/api/v1/test")
def read_root():
    return {"Hello": "World"}

@app.post("/api/v1/summarize")
async def summarize(doc: Document):
    model, tokenizer = load_model()
    print(model)
    print(tokenizer)
    return {
        "document": {
            "title": doc.title,
            "content": doc.content,
        },
    }
