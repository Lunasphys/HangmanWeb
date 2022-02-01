function buttonclick()
{
    var menuList = document.getElementById("choixtxt");
    if (menuList.className == "menuOff")
    {

        menuList.className = "menuOn";
    } else
    {

        menuList.className = "menuOff";
    }
}


const log = document.getElementById('log');

document.addEventListener('keydown', logKey);

function logKey(e) {
  console.log(e.code);
}


