# netflix
With this CLI Tool, Netflix movie list information is instantly retrieved by category group. Movie List information is printed on the screen by Parse instantly from Netflix.com. You can Get Category Group, Movie Count and "Category Group of 5" with "-c 3, -r 6, -f" flags.<br><br>
&#x1F34E;<I>This is Netflix crawler CLI. In every request, It parses Netflix.com. If Netflix.com doesn't response, this service can't work.</I>

![NetflixGithub](https://user-images.githubusercontent.com/9459881/165043647-7e6bad74-623d-405e-8f36-676dd32caf58.png)<br>

<b>Flags:</b>
<table><tr><td>:tv:</td><td><I>netflix</I></td></tr><tr><td></td><td><I>netflix -c 2 -r 6 -f</T></td></tr></table>
<ul>
  <li> -c, --category int   Category Group to be pulled from Netflix: '-f'(group of 5) if there is "-f" flag, -c flag gets a value between 1 to 3, if not 1 and 5 (default 1)</li>
  <li> -f, --five           Number of categories => 5 is assumed instead of 3..</li>
  <li> -h, --help           Help for netflix-cli</li>
  <li> -r, --rowcount int   Total number of movies to be shot by category. Between 1 to 10 (default 5)</li>
</ul>

<b>Default Command:</b> "netflix" </br></br>
<b>Example Usage:</b>

<b>Usage:</b>
  netflix [flags]
<ul>
  <li>"netflix -c 2 -r 6"</li>
  <li>"netflix -r 7 -f"</li>
  <li>"netflix --category 3 --rowcount 8 --five"</li>
  <li>"netflix" => default: 'netflix -c 1 -r 5'</li>
</ul>
**********************************************************************************************************</br>
<i>-c flag gets Max: 3, if there is "-f" flag. If there is no -f flag, -c flag gets Max: 3 value.<br> 
-r flag gets Max:10 value.</i><br> <br>

<b><u>How to install Netflix Cli:</u></b><br>

<table><tr><td><img src="https://user-images.githubusercontent.com/9459881/165053981-38543faf-4bae-4500-8c28-fd5f497e0f46.gif"></img></td>
  <td><i>go install github.com/borakasmer/netflix@latest</i></td></tr></table>

<span style="color: red"><b>&#x1F534;Important:</b></span> You need Go program package to install Netflix-Cli => <a href="https://go.dev/dl/" target="_blank">Go Downloads</a> </br>
<ol>
  <li>"go env" With the command "GOPATH" and "GOROOT" folders of GO can be seen.</li>
  <li>On <b>Mac</b> after "go install ..." the "go/bin/netflix" file under GOPATH => should be copied under "go/bin/" folder under GOROOT.</li>
  <li>In <b>Windows</b>, there is no need to further action :white_check_mark:</li>
</ol>
<img width="427" alt="FRJfHssXEAAqNsP" src="https://user-images.githubusercontent.com/9459881/165074359-572ca085-b1bd-4dbc-840f-43b1690a6319.png">
<b>:green_book: Extra Detail</b><br>
--------------------------------------------------------------------------------------------------------------
<ul>
  <li> "-c" Number of Categories to list. The default value is 1. The first 3 category Groups are listed</li>
  <li> "-c 2" => category group 3 to 6 [3 - 6]</li>
  <li> "-c 3 -f" => Group of categories from 10 to 15 [10 - 15]. (-f => means group of 5)</li>
  <li> "-r" => Number of movies to list. The default value is 5. The first 5 films are printed on the screen</li>
  <li> "netflix" command => By default it means "netflix --category 1 --rowcount 5".</li>
<ul>
    
