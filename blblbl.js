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