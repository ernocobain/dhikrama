

function Toggle(target, target2, result) {
    const main = document.querySelector(target);
    const execu = document.querySelector(target2);
    main.addEventListener('click', () => (
        execu.classList.toggle(result)
    ))
}

function removeToggle(target, target2, result) {
    const main = document.querySelector(target);
    const execu = document.querySelector(target2);
    main.addEventListener('click', () => (
        execu.classList.remove(result)
    ))
}

function addToggle(target, target2, result) {
    const main = document.querySelector(target);
    const execu = document.querySelector(target2);
    main.addEventListener('click', () => (
        execu.classList.add(result)
    ))
}



function ButtonMenusBurger() {
    Toggle(".burger", 'nav', "hidden");
    Toggle(".burger", 'nav', 'z-10');
    Toggle(".burger", 'main', "h-screen");
    Toggle(".burger", 'body', 'bg-gray-500');
    Toggle(".burger", 'body', 'overflow-y-hidden');

    //  remove button

    addToggle("main", 'nav', 'hidden');
    removeToggle("main", 'body', 'overflow-y-hidden');
    removeToggle("main", 'nav', 'z-10');
    removeToggle("main", 'main', "h-screen");
    removeToggle("main", 'body', 'bg-gray-500');
}

export { ButtonMenusBurger }