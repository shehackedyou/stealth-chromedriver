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


## Development
This is an open, transparent, and ideally a community project. Anyone is welcome
to create issues, or submit pull requests.

All pull requests will be taken seriously, and you will not be dismissed or
condescended to.

More details on development are located in the hidden file ".DEVELOPMENT.md". 

### Short Term 
Improve the binary patch functionality and get very basic web "crawler" to run
basic tests. 

Then leverage ideally **"rod"**, or **"colly"**; both are existing well 
developed, high-level crawlers.

It is important to hand off this functionality at least for now so we can focus
on what makes this software unique: evading detection of automated browsing.

At the last step, behavior-based concealment of browser automation, this can be
re-evaluated; determining if it is suitable for our requirements and is using
methods which effectively let us hide close enough to actual browser users that
it becomes too difficult to de-tangle our browser automation from actual usage.

### Long Term
It may feel like an endless *Cat-N-Mouse Game*, but the truth is, the goal is
just to get close enough to a "normal" browser that to block our browsers would
require blocking normal users. This is not an endless arm-race, there is a point
where if we modify the browser and our interaction with it, that hide among the
crowd, in the same way Tor uses **TLS** to hide among normal web traffic.

After generating realistic randomized fingerprints, concealing developer or 
automation tools, the last step will be behavior based on how the input is 
provided explicitly to the JavaScript tools used to identify "normal".

A transparent proxy to block specific connections, file types, and cache locally
common JavaScript, CSS, and other assets will let us obtain further performance
improvements while giving us another layer which we can modify incoming and
outgoing traffic, giving us a second front which we can use to conceal our
automation.

## My Goal
I'm working on a tool to improve upon the great browser extension "Captcha
Buster"; using local ML tools like **OpenCV** to solve **ReCaptcha**, **OCR** 
to solve many **Image+Text Captchas**, and introduce solutions for the other 
new captchas which would not be difficult to automate.

But this requires getting this project up to the behavior based tactics. 

## License 
This project like all projects on this account is GPLv3 unless explicitly stated
otherwise. 


