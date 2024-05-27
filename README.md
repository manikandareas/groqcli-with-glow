# Groq CLI with Glow

<video autoplay controls poster="https://res.cloudinary.com/dzfw66khj/image/upload/v1716823323/d7tckf6iw67mvzw2od5d.png">
  <source src="https://res.cloudinary.com/dzfw66khj/video/upload/v1716820661/aw7mi5f9yzrqxawnxzzx.mov" type="video/mp4">
  <img src="https://res.cloudinary.com/dzfw66khj/image/upload/v1716823323/d7tckf6iw67mvzw2od5d.png" alt="Groq CLI with Glow">
</video>

Groq CLI with Glow is a command-line interface integrated with GroqCloud that allows you to prompt various questions in your terminal. To get started, follow the steps below.

## Prerequisites

  • Groq API Key
  • Groq CLI from this repository
  • Glow, Check this repository for installations glow *[Glow Guide](https://github.com/charmbracelet/glow)*

### Installation

1. Export your Groq API Key using the following command in your terminal:

    ```bash
    export GROQ_API_KEY=YOUR_GROQ_API_KEY
    ```

2. Move the `groq` binary to the `/usr/local/bin` directory using the following command:

    ```bash
    sudo mv groq /usr/local/bin
    ```

### Usage

To use Groq CLI with Glow, simply enter `groq` in your terminal, followed by the prompt:

    groq this is yout prompt

This will launch the Groq CLI with Glow and serve your prompt.

### Flags

To select a specific AI model, use the following flag:

    groq -model <model_name>

Replace `<model_name>` with the desired AI model name.

### Examples

Here's an example of using Groq CLI with Glow to perform a simple query:

    groq can you explain and give an example about grpc with golang

You can also use the `-model` flag to select a specific AI model, like so:

    groq -model=llama3_70b give me an example of tensorflow in python

### Uninstallation

To uninstall Groq CLI with Glow, simply delete the `groq` binary from the `/usr/local/bin` directory.

    sudo rm /usr/local/bin/groq