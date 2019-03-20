var app = new Vue({
  el: '#app',
  data: {
	coords: [],
    message: 'Hello Vue!',
	displayedProds: 1,
	items: [
	{ text: 'Learn JS', 
	color: {backgroundColor: '#009900'}},
	{ text: 'Learn Vue',
	color: {backgroundColor: '#e62e00'}},
	{ text: 'Learn smt',
	color: {backgroundColor: '#ffff00' }}
  ]
  },
  created: function() {
	  this.getCoords();
  },
  methods: {
	  updateProds: function() {
		  this.displayedProds += 1;
		  this.getCoords();
	  },
	  
	  getCoords: function() {
	  var coordinatesArray = [];
		var blocks = document.getElementsByClassName('bl')
		for(var i = 0; i<= blocks.length; i++) {
			var startCoords = 500 + screen.height * 0.16 +  (i * 1183);
			/*console.log(startCoords)*/
			coordinatesArray.push(startCoords);
		}
		this.coords =  coordinatesArray;
  },
  }
})

var Color = net.brehaut.Color

var hd = new Vue({
  el: '#header',
  data: {
    colorQuery: '',
    color: {
      red: 0,
      green: 0,
      blue: 77,
      alpha: 1
    },
    tweenedColor: {}
  },
  created: function () {
    this.tweenedColor = Object.assign({}, this.color)
  },
  watch: {
    color: function () {
      function animate () {
        if (TWEEN.update()) {
          requestAnimationFrame(animate)
        }
      }

      new TWEEN.Tween(this.tweenedColor)
        .to(this.color, 500)
        .start()

      animate()
    }
  },
  computed: {
    tweenedCSSColor: function () {
      return new Color({
        red: this.tweenedColor.red,
        green: this.tweenedColor.green,
        blue: this.tweenedColor.blue,
        alpha: this.tweenedColor.alpha
      }).toCSS()
    }
  },
  methods: {
    updateColor: function () {
      this.color = new Color(this.colorQuery).toRGB()
      this.colorQuery = ''
    }
  }
})

function observeHeaders() {
	var blocks = document.getElementsByClassName('bl')
	for(var i = 0; i< blocks.length; i++) {
		/*console.log(i + ' ' + blocks[i].offsetTop);
		console.log('n' + i + ' ' + app.coords[i]);*/
		if(blocks[i].offsetTop > app.coords[i] && blocks[i].offsetTop < app.coords[i] + 1000 - screen.height * 0.4) {
			hd.colorQuery = blocks[i].style.backgroundColor;
			hd.updateColor();
		}
	}
}

let last_known_scroll_position = 0;
let ticking = false;

window.addEventListener('scroll', Event);

function Event(e) {
  last_known_scroll_position = window.scrollY;

  if (!ticking) {
    window.requestAnimationFrame(function() {
      observeHeaders();
      ticking = false;
    });

    ticking = true;
  }
}