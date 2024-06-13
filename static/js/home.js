
const postContainer = document.querySelector("#container");
const next = document.querySelector("button#next");
const previous = document.querySelector("button#previous");


const fetchPosts = async () => {
    try {
        const urlParams = new URLSearchParams(window.location.search);
        const page = urlParams.get('page');
        const data = await fetch("/data/data-map.json").then(res => res.json());
    
        let index = 0;
        if (!!!Number(page)) {
            index = 0;
        } else {
            index = Number(page);
        };

        const val = await fetch(`/data/${data[index]}`).then(res => res.json());

        renderPosts(val);
        
    } catch (err) {
        console.error(err);
    }
};

const renderPosts = (data = []) => {
    const wrapper = document.createElement("div");
    wrapper.classList.add("wrapper");
    data.forEach((e) => {
       wrapper.appendChild(card(e));
    });
    postContainer.appendChild(wrapper);
};


fetchPosts();


next.addEventListener("click", (e) => {
    const urlParams = new URLSearchParams(window.location.search);
    const page = urlParams.get('page');
    console.log(page);
    if (!!!page) {
        urlParams.set("page", 1);
    } else {
        console.log(page);
        urlParams.set("page", Number(page)+1);
    };
    window.location.search = urlParams.toString();
});

previous.addEventListener("click", (e) => {
    const urlParams = new URLSearchParams(window.location.search);
    const page = urlParams.get('page');
    console.log(page);
    if (!!!page) {
        urlParams.set("page", 1);
    } else {
        console.log(page);
        urlParams.set("page", Number(page)-1);
    };
    window.location.search = urlParams.toString();
});

