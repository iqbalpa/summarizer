from transformers import AutoTokenizer, AutoModelForSeq2SeqLM

def load_model():
    tokenizer = AutoTokenizer.from_pretrained("shorecode/t5-efficient-tiny-summarizer-general-purpose-v2")
    model = AutoModelForSeq2SeqLM.from_pretrained("shorecode/t5-efficient-tiny-summarizer-general-purpose-v2")
    return model, tokenizer

if __name__ == "__main__":
    model, tokenizer = load_model()
    print(f"Model:\n {model}")
    print(f"Tokenizer:\n {tokenizer}")
