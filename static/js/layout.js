(function (window, document) {

    // Hamburg.
    var layout   = document.getElementById('layout'),
        menu     = document.getElementById('menu'),
        menuLink = document.getElementById('menuLink'),
        content  = document.getElementById('main');

    // Toggle menjus.
    let toggleMenus = document.getElementsByClassName("toggle-menu");
    // for (let i = 0; i < toggleMenus.length; i++) {
    //     // toggleMenus[i].get
    // }
    console.log(toggleMenus);

    // Toggle class.
    function toggleClass(element, className) {
        var classes = element.className.split(/\s+/),
            length = classes.length,
            i = 0;

        for(; i < length; i++) {
          if (classes[i] === className) {
            classes.splice(i, 1);
            break;
          }
        }
        // The className is not found
        if (length === classes.length) {
            classes.push(className);
        }

        element.className = classes.join(' ');
    }

    // Toggle all.
    function toggleAll(e) {
        var active = 'active';

        e.preventDefault();
        toggleClass(layout, active);
        toggleClass(menu, active);
        toggleClass(menuLink, active);
    }
    // Show / hide menu.
    menuLink.onclick = function (e) {
        toggleAll(e);
    };
    // Hide menu.
    content.onclick = function(e) {
        if (menu.className.indexOf('active') !== -1) {
            toggleAll(e);
        }
    };

}(this, this.document));