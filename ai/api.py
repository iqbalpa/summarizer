from model import Document
from utils import load_model, summarize
from fastapi import FastAPI

app = FastAPI()

@app.get("/api/v1/test")
def read_root():
    return {"Hello": "World"}

@app.post("/api/v1/summarize")
async def summary(doc: Document):
    model, tokenizer = load_model()
    res = summarize(model, tokenizer, doc.content)
    return {
        "document": {
            "title": doc.title,
            "content": doc.content,
            "summary": res
        },
    }
