var app = new Vue({
  el: '#app',
  data: {
	coords: [],
    message: 'Hello Vue!',
	displayedProds: 2,
	comment: [], 
	items: [
	
	{ text: 'Product 1',
	processor: 'Intel Core i5',
	videocard: 'Nvidia 1010',
	ssd: '128',
	ram: 16,
	price: 276,
	color: {backgroundColor: '#009900'}},
	
	{ text: 'product 2',
	processor: 'Intel Core i7',
	videocard: 'Nvidia 1020',
	ssd: '512',
	ram: 32,
	price: 1000,
	color: {backgroundColor: '#e62e00'}},
	
	{ text: 'product 3',
	processor: 'Intel Core i3',
	videocard: 'Nvidia 1000',
	ssd: '256',
	ram: 8,
	price: 126,
	color: {backgroundColor: '#ffff00' }}
	
  ]
  },
  created: function() {
	  this.getCoords();
  },
  methods: {
	  openCloseCommentBlock: function(index){
		  $(".cmmnt-new").each(function(i) {
			  if(i == index){
				$(this).animate({
					height:'toggle',
				});  
			  }
	});
	  },
	  updateProds: function() {
		  this.displayedProds ++;
		  this.getCoords();
	  },
	  
	  getCoords: function() {
	  var coordinatesArray = [];
		var blocks = document.getElementsByClassName('bl')
		for(var i = 0; i<= blocks.length; i++) {
			var startCoords = 500 + screen.height * 0.27 + (i * ( 1000 + screen.height * 0.25 ) );
			/*console.log(startCoords)*/
			coordinatesArray.push(startCoords);
		}
		this.coords =  coordinatesArray;
	  },
	  sendComment: function (index) {
		  if(this.comment[index] != null &&  this.comment[index] != ''){
			console.log(index + ' ' + this.comment[index]);
			this.comment[index] = '';
			document.getElementsByClassName("form-control").value= '';
		  }
		  this.openCloseCommentBlock(index);
	  },
  }
})

var Color = net.brehaut.Color

var hd = new Vue({
  el: '#header',
  data: {
	  user: {
		  login: '',
	  },
	  
    colorQuery: '',
    color: {
      red: 0,
      green: 0,
      blue: 0,
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
	  scrollToTop: function() {
		 $("html, body").animate({ scrollTop: 0 }, "slow");
	  },
	  
	  testa: function() {
		  this.user.login='aaaa';
		  console.log(app.items[0]);
	  },
	  
	  opemCloseAuthorizationForm: function() {
		  if(document.getElementById("authorizationForm").style.display == "block"){
			  document.getElementById("authorizationForm").style.display = "none";
		  }
		  else {
			  document.getElementById("authorizationForm").style.display="block";
		  }
	  },
    updateColor: function () {
      this.color = new Color(this.colorQuery).toRGB()
      this.colorQuery = ''
    }
  }
})

var af = new Vue({
	el: '#authorizationForm',
	methods:{
		openRegistrationForm: function(){
			document.getElementById("registrationForm").style.display = "block";
			document.getElementById("authorizationForm").style.display = "none";
	  },
	},
})

function observeHeaders() {
	var blocks = document.getElementsByClassName('bl')
	for(var i = 0; i< blocks.length; i++) {
		/*console.log(i + ' ' + blocks[i].offsetTop);
		console.log('S' + i + ' ' + app.coords[i]);*/
		if(blocks[i].offsetTop > app.coords[i]) {
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

var regModal = document.getElementById('registrationForm');

window.onclick = function(event) {
  if (event.target == regModal) {
    regModal.style.display = "none";
  }
}

/* jQuery */

