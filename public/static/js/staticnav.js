
function navStatic(target, [value1, value2, value3], [{optional, optional2, optional3}]) {
    const main = document.querySelector(target);
    window.onscroll = function () {
        if (document.body.scrollTop >= 300 || document.documentElement.scrollTop >= 300) {
            main.classList.add(value1, value2, value3);
        } else {
            main.classList.remove(value1, value2, value3);
        }
    }
}

export { navStatic,}