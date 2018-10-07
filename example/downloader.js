const puppeteer = require('puppeteer');
const http = require('http');
const querystring = require('querystring');
var server = new http.Server();
var browser;

const errReq = {
    'message':'不允许使用非POST请求'
}

const badReq = {
    'message':'Bad Request'
}

const downloadErrReq = {
    'message':'download failed'
}

const uncatchError = {
    'message':'unhandled error occour'
}

server.on('close',function(){
    logInfo('服务器关闭')
})

function logInfo(desc,value) {
    console.log("[Server] [Info]",desc || '',value || '')
}

function logError(desc,value) {
    console.log("[Server] [Error]",desc || '',value || '')
}

server.on('request',function(req,res){    
    res.setTimeout(3600000);
    if (req.method != 'POST') {
        res.writeHead(404, {'Content-Type': 'application/json; charset=utf8'});
        res.write(JSON.stringify(errReq))
        res.end()
        return
    }
    req.setEncoding('utf-8');
    var body = ""
    req.on('data', function (chunk) {
        body += chunk
    });
    req.on('end', function () {
        var dataObject = querystring.parse(body);        
        proxyUrl = dataObject["proxy"] || '';               //代理字符串
        targetUrl = dataObject["target"] || '';             //如果有target     
        waitFor = dataObject["waitfor"] || '';
        logInfo("[proxy]:",proxyUrl + ' [target]:'+targetUrl + "[waitfor]:"+waitFor)
        Run(res,proxyUrl,targetUrl,waitFor)
    });
})

server.listen(47002,function(){      
    server.setTimeout(60000)   
    logInfo('started on :47002')
})

async function InitPuppeteerBrowser(proxyUrl) {        
    opts = [
        '--no-sandbox',
        '--disable-setuid-sandbox',
    ]
    if(proxyUrl && proxyUrl !== ''){
        opts.push('--proxy-server='+ proxyUrl)            
    }
    browser = await puppeteer.launch({
        args: opts,
        timeout:60000,
        ignoreHTTPSErrors: true
    });
}

function ATrim(x) {
    return x.replace(/^\s+|\s+$/gm,'');
}

async function timeout(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function removeHTMLTag(str) {
    str = str.replace(/<\/?[^>]*>/g,''); 
    str = str.replace(/[ | ]*\n/g,'\n'); 
    str=str.replace(/&nbsp;/ig,'');
    str=str.replace(/\s/g,' ');
    return str;
}

async function Run(res,proxyUrl,targetUrl,waitFor) {

    var result = ""
    
    if (proxyUrl.lastIndexOf('/') > 0) {            
        proxyUrl = proxyUrl.substring(proxyUrl.lastIndexOf('/')+1)
        logInfo(proxyUrl.substring(proxyUrl.lastIndexOf('/')))
    }

    //验证代理密码
    proxyInfo = proxyUrl.split("@")
    proxyHost = proxyUrl  
    proxyAuth = ''
    if (proxyInfo.length == 2) {
        proxyHost = proxyInfo[1]  
        proxyAuth = proxyInfo[0]
    }

    await InitPuppeteerBrowser(proxyHost)
    
    const page = await browser.newPage();
    page.setDefaultNavigationTimeout(600000)
    page.waitForNavigation({timeout:600000})
    page.setViewport({
        width: 1920,
        height: 1080,
    }); 

    await page.setExtraHTTPHeaders({
        'Proxy-Authorization': 'Basic ' + Buffer.from(proxyAuth).toString('base64'),
    });
    
    try{
        await page.goto(targetUrl)
        if (waitFor != '') {
            await page.waitfor(waitFor)
        } 
        let bodyHTML = await page.evaluate(() => document.documentElement.innerHTML)
        result = bodyHTML

    }catch(e){
        await browser.close()
        logError(e)
        res.writeHead(200, {'Content-Type': 'text/html; charset=utf8'});        
        res.write(result)
        logInfo('[Server] write data to response body')
        res.end()
        return
    }
    await browser.close()
    res.writeHead(200, {'Content-Type': 'text/html; charset=utf8'});
    res.write(JSON.stringify(result))
    logInfo('[Server] write data to response body')
    res.end()
}

//docker run --rm --shm-size 1G --name puppeteer_downloader -v /home/ajun/gocode/src/spider-libs/basic_downloader/downloader.js:/app/index.js --privileged=true -p 48000:15400 alekzonder/puppeteer