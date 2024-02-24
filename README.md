# Stealth Webdriver (Chromedriver) 
This started out as a port of the popular python project "undetectable chrome
driver", originally an experiment to better understand what strategies were
being employed to avoid detection. In addition, using a compiled language,
receiving a large performance boost. 

Right now it includes some basic crawling functionality and the basic patching
functionality. My plan is to leverage existing crawling projects like "colly",
or "rod", or "playwright"; this would let me focus on the portion that makes
this project unique which is avoiding detection. 

After I reach the limitations of making binary patches to the **chromedriver**,
the next step would be making direct changes to the **cpp** source for the spec
chrome webdriver client. However, now there exists at least a **rust**
alternative available; and ideally these alternatives should be able to talk to
more than one browser since the webdriver protocol is intended to be
cross-browser. 

Which may mean after a single webdriver implementation is determined, either
binary or source patches will need to be made against the browsers themselves.
Because the browser developers have incentive to not let you write automated
browser software that is undetectable, and are almost certainly leaving in
subtle flags that get raised when you use *cdp* or *webdriver v2*. 

## WebDriver
WebDriver is intended to be single protocol for interacting with every browser;
and it is supported by Chrome, Firefox, Edge, IE, and Safari. 

So shifting focus from specifically chromedriver to a more generic webdriver
implementation makes sense as the next step forward for this project. And
get more control over what we can change, to better conceal how our browsers are
being driven; we will want to move from binary patches, to source code patches
and compiling from source. 

For more information, I recommend reviewing: 

  * [WebDriver v2 Documentation](https://www.selenium.dev/documentation/webdriver/browsers/)
  * [Supported Browsers](https://www.selenium.dev/documentation/webdriver/browsers/)


## Long Term
It may feel like an endless cat-n-mouse game, but the truth is, the goal is just
to get close enough to a "normal" browser that to block our browsers would
require blocking normal users. This is not an endless arm-race, there is a point
where if we modify the browser and our interaction with it, that hide among the
crowd, in the same way Tor uses TLS to hide among normal web traffic. 

After we develop realistic fingerprint generation, concealing
developer/automation tools, the last step will be behavior based on how the
input is provided explicitly to the JavaScript tools used to identify "normal". 

A transparent proxy to block specific connections, file types, and cache locally
common JavaScript, CSS, and other assets will let us obtain further performance
improvements while giving us another layer which we can modify incoming and
outgoing traffic, giving us a second front which we can use to conceal our
automation. 

## My Goal
I'm working on a tool to improve upon the great browser extension "captcha
buster"; using local ML tools like OpenCV to solve recaptcha, OCR to solve many
image+text based captchas, and introduce solutions for the other new captchas
which would not be difficult to automate. But this requires getting this project
up to the behavior based tactics. 


