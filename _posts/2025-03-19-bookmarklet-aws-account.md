---
permalink: "bookmarklet-aws-account"
title: "Bookmarklet for translating AWS account numbers to actual accounts"
summary: "Replacing HTML with a hardcoded list of replacements actually is useful"
date: "2025-03-19 22:20"
---

I'm notoriously bad at remembering IP addresses, accounts, things like that, in my head. One thing that's recurring all over our systems at $WORK is AWS account numbers. They are referred to in logs, AWS' own interfaces, monitoring systems, docs/wikis etc. Since it's all web interfaces, we can create a simple solution for that!

This little bookmarklet crudely _replaces an account number with its hardcoded name_. This all happens locally, no risk of leakage. Hopefully, your AWS accounts doesn't change too often so that maintaining this hardcoded list of account numbers ↔ names, becomes a hassle.

Before clicking the bookmarklet:

![AWS account selection menu, listing account numbers](/assets/bookmarklet_aws_before.png)

And after:

![AWS account selection menu, having replaced account numbers with account names](/assets/bookmarklet_aws_after.png)

Here's the snippet:

```javascript
javascript:(function() {
    const bodyClass = "__awsified";
    if (document.body.classList.contains(bodyClass)) {
        return;
    }
    document.body.classList.add(bodyClass);

    const accounts = {
        "123445945445": "SWIMSUIT_FACTORY_PROD",
        "456454354545": "SWIMSUIT_FACTORY_TEST",
        "425642694242": "ZERO_COOL",
        "312857138111": "THE_PLAGUE",
    };

    const ids = Object.keys(accounts).join('|');
    const re = new RegExp(`\\b(${ids})\\b`, 'g');

    const spanStyle = 'font-weight: bold; background-color: yellow; color: black;';

    function wrapMatchesInTextNode(textNode) {
        const text = textNode.textContent;
        const matches = [...text.matchAll(re)];
        if (matches.length === 0) {
            return;
        }

        const frag = document.createDocumentFragment();
        let lastIndex = 0;

        matches.forEach(match => {
            const [matchedText] = match;
            const index = match.index;

            if (index > lastIndex) {
                frag.appendChild(document.createTextNode(text.slice(lastIndex, index)));
            }

            const span = document.createElement('span');
            span.textContent = accounts[matchedText];
            span.setAttribute('title', `Account ID: ${matchedText}`);
            if (document.body.classList.contains(bodyClass)) {
                span.setAttribute('style', spanStyle);
            }

            frag.appendChild(span);
            lastIndex = index + matchedText.length;
        });

        if (lastIndex < text.length) {
            frag.appendChild(document.createTextNode(text.slice(lastIndex)));
        }

        textNode.parentNode.replaceChild(frag, textNode);
    }

    function walkDOM(node) {
        if (node.nodeType === Node.TEXT_NODE) {
            wrapMatchesInTextNode(node);
        } else if (node.nodeType === Node.ELEMENT_NODE) {
            const tag = node.tagName.toLowerCase();
            if (['script', 'style', 'textarea', 'input', 'select'].includes(tag)) {
                return;
            }

            for (let i = 0; i < node.childNodes.length; i++) {
                walkDOM(node.childNodes[i]);
            }
        }
    }

    walkDOM(document.body);
})();
```
