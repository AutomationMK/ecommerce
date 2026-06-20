//////////////////////////////////////
// Custom Navigation Dropdown
//////////////////////////////////////
export default class NavDropdown {
    container;
    header;
    navList;
    dropIcon;

    constructor(container) {
        this.container = container;
        this.header = this.container.querySelector('#nav-header');
        this.dropIcon = this.container.querySelector('#nav-icon');
        this.navList = this.container.querySelector('#nav-list');

        this.init();
    }

    init() {
        // Toggle dropdown on button click
        this.container.addEventListener('mouseover', (e) => {
            e.stopPropagation();
            this.openDropdown();
        });

        // Close dropdown when clicking outside
        this.container.addEventListener('mouseout', (e) => {
            e.stopPropagation();
            this.closeDropdown();
        });
    }

    openDropdown() {
        this.header.classList.remove('text-white');
        this.header.classList.add('text-gray-300');
        this.dropIcon.classList.remove('rotate-0');
        this.dropIcon.classList.remove('text-white');
        this.dropIcon.classList.add('rotate-90');
        this.dropIcon.classList.add('text-gray-300')
        this.navList.classList.remove('border-t-0');
        this.navList.classList.add('border-t-2');
        this.navList.classList.remove('max-h-0');
        this.navList.classList.add('max-h-fit');
        this.navList.classList.add('px-2');
        this.navList.classList.add('py-4');
        this.navList.classList.add('lg:px-4');
        this.navList.classList.add('lg:py-6');
        this.navList.setAttribute('aria-expanded', 'true');
    }

    closeDropdown() {
        this.header.classList.remove('text-gray-300');
        this.header.classList.add('text-white');
        this.dropIcon.classList.remove('rotate-90');
        this.dropIcon.classList.remove('text-gray-300')
        this.dropIcon.classList.add('rotate-0');
        this.dropIcon.classList.add('text-white');
        this.navList.classList.remove('border-t-2');
        this.navList.classList.add('border-t-0');
        this.navList.classList.remove('max-h-fit');
        this.navList.classList.remove('px-2');
        this.navList.classList.remove('py-4');
        this.navList.classList.remove('lg:px-4');
        this.navList.classList.remove('lg:py-6');
        this.navList.classList.add('max-h-0');
        this.navList.setAttribute('aria-expanded', 'false');
    }
}
