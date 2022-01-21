const burgerMenu = document.getElementById('burger');
const navLinks  = document.getElementsByClassName('nav-links')[0];
burgerMenu.addEventListener("click",()=>{
    if(navLinks.style.display === 'none'){
        navLinks.style.display = 'block';
    }else{
        navLinks.style.display = 'none';
    }
})