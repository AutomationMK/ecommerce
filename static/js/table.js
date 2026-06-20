//////////////////////////////////////
// Table
//////////////////////////////////////
"use strict";
export default class Table {
    table;
    body;
    curSlide;
    maxSlide;
    xDown;
    slides;
    btnLeft;
    btnRight;
    dotContainer;
    tableRows;
    activeRows;

    constructor(table) {
        this.curSlide = 0;
        this.table = table;
        this.body = this.table.querySelector(".table-body");
        this.btnLeft = this.table.querySelector(".slide_btn-left");
        this.btnRight = this.table.querySelector(".slide_btn-right");
        this.dotContainer = this.table.querySelector(".dots");
        this.tableRows = Array.from(this.table.querySelectorAll(".table-row"));
        this.activeRows = Array.from(this.tableRows);
        this.maxSlide = 0;

        // implement touch swiping of slider
        this.xDown = null;

        // run button click listeners for slider buttons
        this.btnRight.addEventListener("click", () => {
            this.nextSlide()
        });
        this.btnLeft.addEventListener("click", () => {
            this.prevSlide();
        });

        // run keypad listeners for sliders
        document.addEventListener("keydown", (e) => {
            if (e.key === "ArrowRight") {
                this.nextSlide();
            } else if (e.key === "ArrowLeft") {
                this.prevSlide();
            }
        });

        // run listener for click on dots
        this.dotContainer.addEventListener("click", (e) => {
            if (e.target.classList.contains("dots_dot")) {
                this.curSlide = Number(e.target.dataset.slide);
            }
            this.goToSlide(this.curSlide);
        });

        this.table.addEventListener("touchstart", () => {
            this.handleTouchStart();
        });
        this.table.addEventListener("touchmove", () => {
            this.handleTouchSlide();
        });
    }

    init(numRows, dateRange, search) {
        this.activeRows = this.filterRows(dateRange, search);
        this.setRows(numRows);

        // initialize the slider
        this.goToSlide(0);
    }

    setRows(count) {
        if (this.slides !== undefined) {
            this.slides.forEach((slide) => {
                let tempParent = slide.parentNode;
                this.tableRows.forEach((row) => {
                    let rowParent = row.parentNode;
                    if (this.slides.includes(rowParent)){
                        tempParent.appendChild(row);
                    }
                });
                slide.remove();
            });
        }

        const rowCount = Object.keys(this.activeRows).length;
        let tableBodyCount = 1;
        if (rowCount > count) {
            tableBodyCount = Math.trunc(rowCount/count);
            if (rowCount%count !== 0) {
                tableBodyCount++;
            }
        }
        const slideEls = new Array(tableBodyCount).fill(0);
        slideEls.forEach((_, i, arr) => {
            let slideEl = document.createElement("div");
            slideEl.classList.add("slide", "w-full", "h-full");
            if (i !== 0) {
                slideEl.classList.add("absolute", "-translate-y-full");
            }
            arr[i] = slideEl;
        });
        this.slides = slideEls;
        this.activeRows.forEach((row, i) => {
            let tempParent = row.parentNode;
            let tableIndex = Math.trunc(i/count);
            tempParent.replaceChild(slideEls[tableIndex], row);
            slideEls[tableIndex].appendChild(row);
        });
        this.maxSlide = this.slides.length - 1;

        this.createDots();
        this.goToSlide(0);
    }

    /**
     * Filter based on results from both date range and search
     * @param {string} dateStr - date range select value string
     * @param {string} searchStr - search string
     * @returns {Array} filteredRows - filtered rows that match
     */
    filterRows(dateStr, searchStr) {
        const dateFilterRows = this.filterRowsByDateRange(dateStr);
        const searchFilterRows = this.filterRowsBySearch(searchStr);

        return dateFilterRows.filter((row) => {
            if (searchFilterRows.includes(row)) {
                row.classList.remove("hidden");
                if (!row.classList.contains("table-row")) {
                    row.classList.add("table-row");
                }
                return true;
            } else {
                row.classList.remove("table-row");
                if (!row.classList.contains("hidden")) {
                    row.classList.add("hidden");
                }
                return false;
            }
        });
    }

    /**
     * Filter out rows by search term
     * @param {string} searchStr - search string
     * @returns {Array} filteredRows - filtered rows that match the search
     */
    filterRowsBySearch(searchStr) {
        // if search is empty do not perform any filter
        if (searchStr === "") {
            return Array.from(this.tableRows);
        } else {
            let inSearch = false;
            let filteredRows = this.tableRows.filter((row) => {
                inSearch = false;
                row.querySelectorAll(".table-column").forEach((col) => {
                    // set inSearch variable if any columns match the search
                    let colStr = String(col.innerText).toLowerCase();
                    if (colStr.includes(searchStr.toLowerCase())) {
                        inSearch = true;
                    }
                });

                // check if any columns where matched with search
                if (inSearch) {
                    return true;
                } else {
                    return false;
                }
            });
            return filteredRows;
        }
    }

    /**
     * Filter out rows not in the selected date range
     * @param {string} selectStr - dateRange select value string
     * @returns {Array} filteredRows - filtered rows from date Range
     */
    filterRowsByDateRange(selectStr) {
        const maxDate = new Date();
        const minDate = new Date();
        if (selectStr === "This Month") {
            minDate.setDate(1);
            minDate.setHours(0,0,0,0);
            this.#incMonth(maxDate);
        } else if (selectStr === "Next Month") {
            this.#incMonth(minDate);
            this.#incMonth(maxDate, 2);
        } else if (selectStr === "3 Months") {
            minDate.setDate(1);
            minDate.setHours(0,0,0,0);
            this.#incMonth(maxDate, 3);
        } else if (selectStr === "6 Months") {
            minDate.setDate(1);
            minDate.setHours(0,0,0,0);
            this.#incMonth(maxDate, 6);
        } else if (selectStr === "1 Year") {
            minDate.setDate(1);
            minDate.setHours(0,0,0,0);
            this.#incMonth(maxDate, 12);
        } else {
            this.tableRows.forEach((row) => {
                row.classList.remove("hidden");
                if (!row.classList.contains("table-row")) {
                    row.classList.add("table-row");
                }
            });
            return Array.from(this.tableRows);
        }

        let filteredRows = this.tableRows.filter((row) => {
            let colDateEl = row.querySelector(".date-column");
            const colDate = new Date(colDateEl.dataset.isodate);

            if (colDate.getTime() >= minDate.getTime() && colDate.getTime() < maxDate.getTime()) {
                row.classList.remove("hidden");
                if (!row.classList.contains("table-row")) {
                    row.classList.add("table-row");
                }
                return true;
            } else {
                row.classList.remove("table-row");
                if (!row.classList.contains("hidden")) {
                    row.classList.add("hidden");
                }
                return false;
            }
        });
        return filteredRows;
    }

    /**
     * Increase a Date object to the beginning of the next month
     * @param {Date} date - date to be increased
     * @param {integer} times - amount of months to increase to
     */
    #incMonth(date, times=1) {
        if (times > 1) {
            for (let i=0; i < times; i++) {
                date.setDate(1);
                date.setHours(0,0,0,0);
                let d = date.getDate();
                date.setMonth(date.getMonth() + 1);
                if (date.getDate() !== d) {
                    date.setDate(0);
                }
            }
        } else {
            date.setDate(1);
            date.setHours(0,0,0,0);
            let d = date.getDate();
            date.setMonth(date.getMonth() + 1);
            if (date.getDate() !== d) {
                date.setDate(0);
            }
        }
    }
    
    // create dot slider elements based on amount of slides
    createDots() {
        this.dotContainer.querySelectorAll(".dots_dot").forEach((dot) => {
            dot.remove();
        })
        if (this.slides.length > 1) {
            this.slides.forEach((_, i) => {
                this.dotContainer.insertAdjacentHTML(
                    "beforeend",
                    `<button class="dots_dot" data-slide="${i}"></button>`,
                );
            });
            this.btnRight.classList.remove("hidden");
            this.btnLeft.classList.remove("hidden");
        } else {
            this.dotContainer.insertAdjacentHTML(
                "beforeend",
                `<button class="dots_dot hidden" data-slide="0"></button>`,
            );
            this.btnRight.classList.add("hidden");
            this.btnLeft.classList.add("hidden");
        }
    }

    // activate a dot to show which slide is active
    activateDot(index) {
        this.dotContainer
            .querySelectorAll(".dots_dot")
            .forEach((dot) => dot.classList.remove("dots_dot-active"));

        this.dotContainer
            .querySelector(`.dots_dot[data-slide="${index}"]`)
            .classList.add("dots_dot-active");
    }

    // move to slide with 0 based indexing
    goToSlide(index) {
        // check if index is out of bounds
        // if so set the global curSlide to other extreme
        if (index > this.maxSlide) this.curSlide = index = 0;
        if (index < 0) this.curSlide = index = this.maxSlide;
        
        this.activateDot(index);
        this.slides.forEach(
            (s, i) => (s.style.transform = `translateX(${100 * (i - index)}%)`),
        );
    }

    // move to next slide
    nextSlide() {
        this.curSlide++;
        this.goToSlide(this.curSlide);
    }

    // move to previous slide
    prevSlide() {
        this.curSlide--;
        this.goToSlide(this.curSlide);
    }

    handleTouchStart(e) {
        const firstTouch = e.touches[0];
        this.xDown = firstTouch.clientX;
    }

    handleTouchSlide(e) {
        if (!this.xDown) {
            return;
        }

        var xUp = e.touches[0].clientX;

        var xDiff = xDown - xUp;

        if (xDiff > 5 && curSlide < maxSlide) {
            // right slide
            nextSlide();
        } else if (xDiff < -5 && curSlide > 0) {
            // left slide
            prevSlide();
        }
        // reset global value
        this.xDown = null;
    }
}
