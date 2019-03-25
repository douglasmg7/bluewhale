(function (window, document) {

    // Hamburg.
    var menu     = document.getElementById('menu');
    var menuToggle = document.getElementById('menu-toggle');

    // Toggle all.
    function toggleMenu(e) {
        e.preventDefault();
        toggleClass(menu, "active");
        toggleClass(menuToggle, "active");
    }

    // Show and hide menu.
    menuToggle.onclick = function (e) {
        toggleMenu(e);
    };

    // Show and hide sub-menus. 
    let toggleItems = document.getElementsByClassName("toggle-items");
    for (let i = 0; i < toggleItems.length; i++) {
        let subMenu = toggleItems[i].nextElementSibling;
        // Show sub-menu.
        toggleItems[i].onclick =  function(e){
            toggleClass(subMenu, "active");
        }
        // Hide sub-menu.
        let backElements = subMenu.getElementsByClassName("back");
        backElements[0].onclick = function(e){
            toggleClass(subMenu, "active");
        }
    }

    // Collapse element.
    function collapseElement(element){
        // Get the height of the element's inner content, regardless of its actual size.
        let elementHeight = element.scrollHeight; 
        // Temporarily disable all css transitions.
        let elementTransition = element.style.transition; 
        element.style.transition = ""
        // On the next frame (as soon as the previous style change has taken effect),
        // explicitly set the element's height to its current pixel height, so we aren't transitioning out of 'auto'.
        requestAnimationFrame(function(){
            element.style.height = elementHeight + "px";
            element.style.transition = elementTransition;
            // On the next frame (as soon as the previous style change has taken effect),
            // have the element transition to height: 0.
            requestAnimationFrame(function(){
                element.style.height = "0" + "px";
                element.setAttribute('data-expanded', "false");
            });
        });
    }
    // Expand element.
    function ExpandElement(element){
        // Get the height of the element's inner content, regardless of its actual size.
        let elementHeight = element.scrollHeight; 
        // Have the element transition to the height of its inner content.
        element.style.height = elementHeight + "px";
        // When the next css transition finishes (which should be the one we just triggered).
        element.addEventListener('transitionend', function(e){
            // Remove this event listener so it only gets triggered once.
            element.removeEventListener('transitionend', arguments.callee);
            // Remove "height" from the element's inline styles, so it can return to its initial value.
            // element.style.height = null;
        });
        element.setAttribute('data-expanded', "true");
    }
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


}(this, this.document));
