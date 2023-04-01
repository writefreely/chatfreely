# ChatFreely

ChatFreely is a super-intelligent, sentient and sweet AI. Just kidding. It's a program that generates brand new blog posts from the corpus of text on your own WriteFreely blog.

It uses a [Markov chain](https://en.wikipedia.org/wiki/Markov_chain) to generate often nonsensical but sometimes amusing new strings of words. It's functionally like ChatGPT, but more silly and without all the non-consensual training on vast amounts of intellectual property.

## Limitations

_For now, it only works with Write.as._

Otherwise, like most modern "AI" tools:

* ChatFreely will confidently poop out language without any concern for truth or reality. It has no capacity for any kind of logical inference that would allow it to understand truth as humans do.
* It will produce instructions and content based on the data it was trained on, for better or worse.
* It has no knowledge of the world and events, in 2023 or in any other year. It does not have knowledge. It is a mathematical model coded into a computer program.

## Getting started

With Go installed, open a terminal and run:

```bash
go get github.com/writefreely/chatfreely
```

### Training

Next train the "AI" on your Write.as blog with the following command, replacing `[blog-alias]` with your own.

The `-o` flag sets the order of your Markov model -- setting it to `2` makes it more coherent; setting it to `1` makes it more unhinged but more original.

```bash
chatfreely train -c [blog-alias] -o 2
```

#### Using with Write.as

Write.as implements rate-limiting on its post-retrieval API that [requires an application key](https://write.as/me/applications). Follow the instructions there to retrieve yours, and then set an environment variable, `WRITEAS_APP_KEY=your-app-key-here...`, before training your model. 

### Generating

Finally, generate a brand new post, again specifying the order (`-o`) that you used to train the model:

```bash
chatfreely gen -c [blog-alias] -o 2
```

## Commands

```
NAME:
   ChatFreely - Generative "AI" that learns from WriteFreely blogs.

USAGE:
   chatfreely [global options] command [command options] [arguments...]

COMMANDS:
   train          Train the markov chain.
   generate, gen  Generate a blog post from training data.
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### train

Train the markov chain.

```
USAGE:
   chatfreely train [command options] [arguments...]

OPTIONS:
   --alias value, -c value  Alias of the WriteFreely collection to train on
   --order value, -o value  Markov chain order (recommend 1 or 2)
   --help, -h               show help
```

### generate

Generate a blog post from training data.

```
USAGE:
   chatfreely generate [command options] [arguments...]

OPTIONS:
   --alias value, -c value  Alias of the WriteFreely collection
   --order value, -o value  Markov chain order (same as training)
   --help, -h               show help
```