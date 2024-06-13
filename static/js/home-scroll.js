const postContainer = document.querySelector("#container");
const targetElement = document.querySelector("#view");
let siteData = [];

let hold = false;
let index = 0;
let stop = false;

const exitScroll = () => {
  window.removeEventListener('scroll', onScroll);
  targetElement.style.height = "0px";
}

const renderPosts = (data = []) => {
    data.forEach((e) => {
       postContainer.appendChild(card(e));
    });
    setTimeout(() => {
        hold = false;
        index++;
    }, 1000);
};


const initData = () => {
    fetch("/data/data-map.json").then(res => res.json()).then(res => {
        siteData = res;
    }).catch(err => {
        console.error(err);
    });
};


  // Function to check if the element is in view
  const isElementInView = (element) => {
    const rect = element.getBoundingClientRect();
    const windowHeight = (window.innerHeight || document.documentElement.clientHeight);
    const windowWidth = (window.innerWidth || document.documentElement.clientWidth);

    // Check if any part of the element is in the viewport
    return (
      rect.top <= windowHeight && 
      rect.bottom >= 0 &&
      rect.left <= windowWidth &&
      rect.right >= 0
    );
  };


const fetchPosts = async () => {
    try {
        if (index == 0) {
          if (siteData.length == 1) {
            exitScroll();
          }
          index++
          return
        }

        hold = true;
        const file = siteData[index]

        if (!!!file) {
          exitScroll();
          return
        }

        const val = await fetch(`/data/${file}`).then(res => res.json());

        renderPosts(val);
        
    } catch (err) {
      exitScroll();
    };
};

  const onScroll = () => {
    if (isElementInView(targetElement) && !hold) {
      console.log('The target element has come into view!');
      fetchPosts();
    };
  };

initData();


document.addEventListener('DOMContentLoaded', () => {


  // Scroll event listener

  // Add the scroll event listener
  window.addEventListener('scroll', onScroll);

  // Check if the element is in view on page load
  onScroll();
});