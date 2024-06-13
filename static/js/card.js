
const tagsTag = (tags = []) => {
    const tagsEle = document.createElement("div");
    tagsEle.classList.add("tags");
    tags.forEach(tag => {
        tagsEle.appendChild(tagTag(tag));
    });
    return tagsEle;
};

const tagTag = (tag) => {
    const tagEle = document.createElement("a");
    tagEle.classList.add("tag");
    tagEle.href = `/tag/${tag}`;
    tagEle.textContent = `#${tag}`;
    return tagEle;
};

const landingImageTag = (src = "", alt = "") => {
    const img = document.createElement("img");
    img.src = src;
    img.alt = alt;
    return img
};


const card = ({ title = "", tags = [], category = "", landingImage = "", date = "" }) => {
    const card = document.createElement("div");
    const h1 = document.createElement("h1");
    const metaData = document.createElement("div");
    const categoryEle = document.createElement("a");
    const dateEle = document.createElement("div");
    const imageLink = document.createElement("a");
    const titleLink = document.createElement("a");


    metaData.classList.add("card__meta-data");
    h1.textContent = title;
    card.classList.add("card");
    categoryEle.classList.add("card__category");
    categoryEle.href = `/category/${category}`;
    imageLink.href = `/category/${category}`;
    titleLink.href = `/category/${category}`;
    categoryEle.textContent = category;
    metaData.appendChild(categoryEle);
    metaData.appendChild(tagsTag(tags));
    dateEle.classList.add("card__date");
    dateEle.textContent = String(date).slice(0,16);
    metaData.appendChild(dateEle);
    imageLink.appendChild(landingImageTag(landingImage));
    titleLink.appendChild(h1);

    
    card.appendChild(imageLink);
    card.appendChild(titleLink);
    card.appendChild(metaData);
    return card;
};
