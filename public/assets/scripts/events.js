document.querySelector("#RingRing").addEventListener("click", e => {
    if(/Mobi|Android/i.test(navigator.userAgent)) 
        return window.location.href= `tel:${e.target.value}`
    CopyToClipboard(e.target.value)
})