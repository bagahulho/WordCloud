# Telegram WordCloud

Hello. This repository will allow you to make a beautiful picture of the words from your chats, which will line up in the shape you want.

#### *Please follow the instructions below and then you will definitely succeed.*

1. The first thing we need to do is to export either one chat or all chats from Telegram.  
   [This shows you how to do it.](https://translated.turbopages.org/proxy_u/en-ru.ru.acef52d3-66be5179-e2c8be42-74722d776562/https/www.thewindowsclub.com/how-to-export-chat-and-group-data-in-telegram)  
   You need to uncheck all the checkboxes, select JSON format, and remember the path to the file. If you want to export not just one chat, but all personal chats, then Settings -> Advanced -> Export Telegram data and also uncheck all the boxes except “personal chats” and select “machine-readable JSON”.

2. Choose an image for the form in PNG format and download it (it should be without a background).  
   [Here’s an example](https://yandex.ru/images/search?from=tabbar&img_url=https%3A%2F%2Fwww.clipartmax.com%2Fpng%2Ffull%2F50-504644_gold-crown-clipart-transparent-background-collection-transparent-background-crown-clipart.png&lr=1092&pos=0&rpt=simage&text=png%20%D0%BA%D0%BE%D1%80%D0%BD%D0%B0).

3. [Install Golang](https://go.dev/doc/install) if you do not have it already.

4. Install the WordCloud CLI tool by running the following command:
   ```bash
   go install github.com/bagahulho/WordCloud@make-cli
   ```
5. Add Go binaries to your PATH:

   If the WordCloud command is not found, you may need to add the Go binaries to your system's PATH variable.

   - macOS/Linux:
      1. Add the Go bin directory to your PATH:
         ```bash
         export PATH=$PATH:$HOME/go/bin
         ```
      2. To make this change permanent, add the above line to your ~/.bashrc, ~/.zshrc, or ~/.profile file.
      3. Apply the changes:
         ```bash
         source ~/.bashrc  # or source ~/.zshrc
         ```

   - Windows:
      1. Go binaries are typically located in %USERPROFILE%\go\bin.
      2. Add this directory to the PATH environment variable:
         - Open "System Properties" > "Environment Variables".
         - Find and edit the Path variable in "User variables".
         - Add a new entry for %USERPROFILE%\go\bin.
      3. Restart your command prompt or terminal for the changes to take effect.

6. Run the tool with the --help flag to see further instructions for use:
   ```bash
   WordCloud --help
   ```

7. Example of use:
   ```bash
   WordCloud makeSingle --json /path/to/your/telegram.json --output /path/to/output/image.png --mask /path/to/your/mask.png
   ```
   Replace /path/to/your/telegram.json with the actual path to your exported Telegram data in JSON format,  
   /path/to/output/image.png with the path where you want to save the generated word cloud image,  
   and /path/to/your/mask.png with the path to the PNG image you want to use as the mask shape for your word cloud.
   Use makeSingle (mS) if you exported one chat, and makeMulti (mM) if you exported all chats.
   You can also specify the size of the resulting image using the --width and --height flags (the larger the size you choose, the more words will fit).

### Example
- mask:

![mask image](https://github.com/bagahulho/WordCloud/blob/make-cli/example/mask.png)

- result:

![result image](https://github.com/bagahulho/WordCloud/blob/make-cli/example/output.png)