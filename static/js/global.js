const menu = document.querySelector("button#menu");
const drawer = document.querySelector("#drawer");
const copied = document.querySelector("#copied")




const toggleMenu = () => {
    const show = drawer.classList.contains("drawer--show");
    if (show) {
        drawer.classList.remove("drawer--show");
        document.body.classList.remove("open-menu")
    } else {
        drawer.classList.add("drawer--show");
        document.body.classList.add("open-menu")
    };
};

const hideCopied = () => {
    setTimeout(() => {
        copied.classList.remove("copied-show")
    }, 3000);
};


const copyURL = () => {
    if (!!window.navigator && !!window.navigator.clipboard) {
        window.navigator.clipboard.writeText(window.location.href).then(res => {
            copied.classList.add("copied-show")
            hideCopied()
        })
    }
}