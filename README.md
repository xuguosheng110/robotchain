# 🤖🔗 RobotChain

⚡ Welcome to the new world of general artificial intelligence and general robotics. ⚡

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 🤔 What is this?

> Large language models (LLMs) are becoming a transformative technology that enables developers to build applications previously unattainable. However, using LLMs in isolation is often insufficient to create a truly powerful application. The real power comes from combining them with other sources of computation or knowledge - **the exploration of their integration with robotics is just beginning**.

> RobotChain is an open-source development framework based on universal robotics and general artificial intelligence technologies, serving as the cornerstone of the integration of universal robotics and general artificial intelligence.

> RobotChain empowers the development and application of universal robotics and general artificial intelligence to become simpler.

- 🧬 RobotChain is developed and built based on `GoLang`, `Vue + Vite + Typescript`, `Python`, and `ROS2`. It runs on the `Ubuntu` operating system.

- ✨ The inspiration for this project comes from [LangChain](https://github.com/hwchase17/langchain).

## 📖 Documentation

> Please see [here](https://geekros.github.io) for full documentation on:

✏ Content editing in progress...

## 🚀 What can this help with?

> Large language models (LLMs) have a wide range of applications, but RobotChain mainly provides support and assistance in these main areas:

✏ Content editing in progress...

## Quick Install

> 🚨 Currently, it only supports running on Ubuntu 20.04+ operating system.

```shell
sudo curl -s https://cdn.geekros.com/robotchain/install.sh|bash
```

> 🪄 Alternatively, you can also deploy using the following method:

```shell
sudo touch /etc/apt/sources.list.d/geekros.list
```

```shell
sudo wget -q -P /usr/share/geekros/ https://cdn.geekros.com/robotchain/pgp-key.public
```

```shell
echo "deb [signed-by=/usr/share/geekros/pgp-key.public] https://ubuntu.geekros.com/ubuntu-ports/ focal main" | sudo tee /etc/apt/sources.list.d/geekros.list
```

```shell
sudo apt -y update && sudo apt -y install robotchain
```

## 🛸 Build

> 🚨 Currently, it only supports running on Ubuntu 20.04+ operating system.

#### ➡ Compilation environment requirements:

- Golang Version ≥ 1.18.0
- Python Version ≥ 3.8.0

#### ➡ Explanation of source code directory structure:

``` python
├─robotchain
│  ├─framework # Coreframework based on Golang
│  ├─package # Common function packages based on ROS2.
│  ├─template # Front-end web UI
```

✏ Content editing in progress...

## 💁 Contributing

> RobotChain is currently in its early stage and does not accept any external contributions at the moment. If you encounter any issues during use, please feel free to submit them in the issues section.