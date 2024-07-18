#!/usr/bin/env python3

# from https://scribe.rip/swlh/how-simple-is-it-to-build-an-end-to-end-item-based-recommender-system-90f6d959e7c2


####### files

import os

files = {}

for filename in os.listdir("_posts"):
    filepath = os.path.join("_posts", filename)
    if os.path.isfile(filepath):
        with open(filepath, "r") as f:
            files[filepath] = {"content": f.read().replace('\n', '')}


####### tf

import tensorflow_hub as hub
import tensorflow_text

embed = hub.load("https://tfhub.dev/google/universal-sentence-encoder-multilingual/3")

for f in files:
    files[f]["vector"] = embed(files[f]["content"])[0]



###### scikit-learn

import numpy as np
from sklearn.metrics.pairwise import cosine_similarity

print([f["content"] for f in files][0])
vectors = [f["vector"] for f in files]
X = np.array(vectors)

cos_sim = cosine_similarity(X)
cos_indices = np.vstack([np.argsort(-arr) for arr in cos_sim])

for i, f in enumerate(files):
    f["cosine"] = cos_indices[i][1:21]


print(next(files))
print(next(files))
