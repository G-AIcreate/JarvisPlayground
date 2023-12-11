from llama_index.llms import LlamaCPP
from llama_index.llms.llama_utils import (
    messages_to_prompt,
    completion_to_prompt,
)
import os
# print(os.getcwd()+"33")
# model_url = "https://huggingface.co/TheBloke/Llama-2-13B-chat-GGUF/blob/main/llama-2-13b-chat.Q5_K_M.gguf"
model_path = "ai-model-training-service/models/GLlamaModel/llama-2-13b-chat.Q4_0.gguf"

llm = LlamaCPP(
    # You can pass in the URL to a GGML model to download it automatically
    model_url=None,
    # optionally, you can set the path to a pre-downloaded model instead of model_url
    model_path=model_path,
    temperature=0.1,
    max_new_tokens=256,
    # llama2 has a context window of 4096 tokens, but we set it lower to allow for some wiggle room
    context_window=3900,
    # kwargs to pass to __call__()
    generate_kwargs={},
    # kwargs to pass to __init__()
    # set to at least 1 to use GPU
    model_kwargs={"n_gpu_layers": 1},
    # transform inputs into Llama2 format
    messages_to_prompt=messages_to_prompt,
    completion_to_prompt=completion_to_prompt,
    verbose=True,
)

# response = llm.complete("Hello! 日本の人口は?")
# print(response.text)

# 応答をストリーミングする場合
def getAnswer(question):
    question = dequeue(question)
    answer = llm.stream_complete(question)
    return answer

def dequeue(question):

    return question


# def getAnswer(question):
#     answer = llm.stream_complete(getQuestion())
#     return answer
def enqueue(answer):
    return
    # ssss

def main():
    return

if __name__=="__main__":
    main()

# response_iter = llm.stream_complete(getQuestion())
# for response in response_iter:
#     print(response.delta, end="", flush=True)