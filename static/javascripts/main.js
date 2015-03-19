(function($){
	
	"use strict";
	
    $(document).ready(function(){

    	$(".error").hide();

		jQuery('#countdown_dashboard').countDown({
				targetDate: {
					'day': 		23, // Put the date here
					'month': 	5, // Month
					'year': 	2015, // Year
					'hour': 	0,
					'min': 		0,
					'sec': 		0
				} //,omitWeeks: true
		});
			
		$(".party-container").on("click", "a", function(e) {
			$("#party-popup").modal();
			var modalHeight = $(window).height() * .6; // Set the height to 60% of the screen size
			$(".modal-body").css("height", modalHeight);
			e.preventDefault();
			e.stopPropagation();
		});

		var intRegex = /^\d+$/;

		var validateRSVPLookup = function() {
			var valid = true;
			// Make sure that RSVP ID provided and a number
			var id = $.trim($("#rsvpId").val());
			if (!!id && intRegex.test(id)) {
				$("#rsvpId-error").hide();
				$("#rsvpId").val(id);
			}
			else {
				$("#rsvpId-error").show();
				valid = false;
			}
			// Make sure first name provided
			var firstName = $.trim($("#first1").val());
			if (!!firstName) {
				$("#first1-error").hide();
				$("#first1").val(firstName);
			}
			else {
				$("#first1-error").show();
				valid = false;
			}
			// Make sure last name provided
			var lastName = $.trim($("#last1").val());
			if (!!lastName) {
				$("#last1-error").hide();
				$("#last1").val(lastName);
			}
			else {
				$("#last1-error").show();
				valid = false;
			}
			return valid;
		};

		var validateSecondPrimary = function() {
			var valid = true;
			// Check that attendance provided for primary
			if ($("#attending2").val() !== "-1") {
				$("#attending2-error").hide();
			}
			else {
				$("#attending2-error").show();
				valid = false;
			}
			// Check that dinner selection has been made for primary
			if ($("#food2").is(":visible")) {
				if ($("#food2").val() !== "-1") {
					$("#food2-error").hide();
				}
				else {
					$("#food2-error").show();
					valid = false;
				}
			}
			return valid;
		};

		var validateGuest = function() {
			var valid = true;
			// Only validate a guest if one was provided
			if (!$("#guestForm").is(":visible")) return valid;
			// Make sure first name provided
			// Make sure first name provided
			var firstName = $.trim($("#guestFirst").val());
			if (!!firstName) {
				$("#gFirst-error").hide();
				$("#guestFirst").val(firstName);
			}
			else {
				$("#gFirst-error").show();
				valid = false;
			}
			// Make sure last name provided
			var lastName = $.trim($("#guestLast").val());
			if (!!lastName) {
				$("#gLast-error").hide();
				$("#guestLast").val(lastName);
			}
			else {
				$("#gLast-error").show();
				valid = false;
			}
			// Check that dinner selection has been made for the guest
			if ($("#guestFood").val() !== "-1") {
				$("#gFood-error").hide();
			}
			else {
				$("#gFood-error").show();
				valid = false;
			}
			return valid;
		};

		var validateRSVPReply = function() {
			var valid = true;
			// Check that attendance provided for primary
			if ($("#attending1").val() !== "-1") {
				$("#attending1-error").hide();
			}
			else {
				$("#attending1-error").show();
				valid = false;
			}
			// Check that dinner selection has been made for primary
			if ($("#food1").is(":visible")) {
				if ($("#food1").val() !== "-1") {
					$("#food1-error").hide();
				}
				else {
					$("#food1-error").show();
					valid = false;
				}
			}
			// Check to see if there's a second primary
			var restValid = true;
			if ($("#first2").length) {
				restValid = validateSecondPrimary();
			}
			else {
				restValid = validateGuest();
			}
			return valid && restValid;
		};

		var getRSVPForm = function(rsvpURL) {
			$.ajax({
				url: rsvpURL,
				type: "GET",
				cache: false,
				timeout: 5000
			}).done(function(data, textStatus, jqXHR) {
				var errHeader = $("#rsvpwrap h3");
				if (errHeader.length) {
					errHeader.remove();
				}
				if (!!jqXHR.responseText) {
					$("#rsvpwrap form").replaceWith(data);
					$(".error").hide();
				}
				else {
					$("#rsvpwrap form").replaceWith("<h3>Doh! It seems the gremlins have done something. The server has vanished. Please come back later and try again. Worst-case snail mail wins.</h3>");
				}
			}).fail(function(jqXHR, textStatus, errorThrown) {
				var errHeader = $("#rsvpwrap h3");
				if (errHeader.length) {
					errHeader.remove();
				}
				if (!!jqXHR.responseText) {
					$("#rsvpwrap form").replaceWith(jqXHR.responseText);
					$(".error").hide();
				}
				else {
					$("#rsvpwrap form").replaceWith("<h3>Doh! It seems the gremlins have done something. The server has vanished. Please come back later and try again. Worst-case snail mail wins.</h3>");
				}
			});
		};

		$("#rsvpwrap").on("change", "select", function(e) {
			var selectBox = $(e.target);
			var selectWidget = selectBox.attr("id");
			var selectValue = selectBox.val();
			if (selectWidget === "attending1") {
				if (selectValue === "1") {
					$("#food1Box").slideDown();
					$("#attending1-error").hide();
				}
				else if (selectValue === "0") {
					$("#food1Box").slideUp();
					$("#food1").val("-1");
					$("#attending1-error").hide();
				}

			}
			else if (selectWidget === "attending2") {
				if (selectValue === "1") {
					$("#food2Box").slideDown();
					$("#attending2-error").hide();
				}
				else if (selectValue === "0") {
					$("#food2Box").slideUp();
					$("#food2").val("-1");
					$("#attending2-error").hide();
				}
			}
			else if (selectWidget === "food1") {
				if (selectValue !== "-1") {
					$("#food1-error").hide();
				}
			}
			else if (selectWidget === "food2") {
				if (selectValue !== "-1") {
					$("#food2-error").hide();
				}
			}
			else if (selectWidget === "guestFood") {
				if (selectValue !== "-1") {
					$("#gFood-error").hide();
				}
			}
		});

		$("#rsvpwrap").on("focus", ".text-input, select", function () {
	        $(this).css({border:"2px solid #de675f"});
	        $(this).css({background:"#fff"});
	    });

	    $("#rsvpwrap").on("blur", ".text-input, select", function () {
	        $(this).css({border:"2px solid #fff"});
	        $(this).css({background:"transparent"});
	    });

	    $("#rsvpwrap").on("click", "#add-guest-btn", function() {
	    	var guestButton = $("#add-guest-btn");
	    	if (guestButton.val() === "Add Guest") {
	    		guestButton.val("Remove Guest");
	    		$("#guestForm").slideDown();
	    	}
	    	else {
	    		$("#guestForm").slideUp();
	    		$("#guestFirst").val("");
	    		$("#guestLast").val("");
	    		$("#guestFood").val("-1");
	    		$("#gFood-error").hide();
	    		$("#gLast-error").hide();
	    		$("#gFirst-error").hide();
	    		guestButton.val("Add Guest");
	    	}
	    	
	    });

		/* Form submission capture */
		$("#rsvpwrap").on("submit", "#start-rsvp", function(e) {
			e.preventDefault();
			var validForm = validateRSVPLookup();
			if (!validForm) return;
			$("#start-rsvp-btn").prop("disabled", true);
			// Submit the form and process response
			var form = $(this);
			var postData = form.serialize().toLowerCase();
			$.ajax({
				url: form.attr("action"),
				type: form.attr("method"),
				cache: false,
				data: postData,
				timeout: 5000
			}).done(function(data, textStatus, jqXHR) {
				var rsvpFormURL = jqXHR.getResponseHeader("Location");
				if (!!rsvpFormURL) {
					getRSVPForm(rsvpFormURL);
				}
				else {
					$("#rsvpwrap form").replaceWith("<h3>Doh! It seems the gremlins have done something. The server has vanished. Please come back later and try again. Worst-case snail mail wins.</h3>");
				}
			}).fail(function(jqXHR, textStatus, errorThrown) {
				var errHeader = $("#rsvpwrap h3");
				if (errHeader.length) {
					errHeader.remove();
				}
				if (!!jqXHR.responseText) {
					$("#rsvpwrap form").replaceWith(jqXHR.responseText);
					$(".error").hide();
				}
				else {
					$("#rsvpwrap form").replaceWith("<h3>Doh! It seems the gremlins have done something. The server has vanished. Please come back later and try again. Worst-case snail mail wins.</h3>");
				}
			});
		});

		$("#rsvpwrap").on("submit", "#rsvp-reply", function(e) {
			e.preventDefault();
			var validForm = validateRSVPReply();
			if (!validForm) return;
			$("#reply-submit-btn").prop("disabled", true);
			var form = $(this);
			var postData = form.serialize().toLowerCase();
			$.ajax({
				url: form.attr("action"),
				type: form.attr("method"),
				cache: false,
				data: postData,
				timeout: 5000
			}).done(function(data, textStatus, jqXHR) {
				var errHeader = $("#rsvpwrap h3");
				if (errHeader.length) {
					errHeader.remove();
				}
				if (jqXHR.status === 200) {
					$("#rsvpwrap form").replaceWith("<h1>Thank you for taking the time to RSVP!</h1>");
					$("a[href='#rsvp']").trigger("click");
				}
				else {
					$("#rsvpwrap form").replaceWith("<h3>Doh! It seems the gremlins have done something. Please come back later and try again. Worst-case snail mail wins.</h3>");
				}
			}).fail(function(jqXHR, textStatus, errorThrown) {
				var errHeader = $("#rsvpwrap h3");
				if (errHeader.length) {
					errHeader.remove();
				}
				if (!!jqXHR.responseText) {
					$("#rsvpwrap form").replaceWith(jqXHR.responseText);
					$(".error").hide();
				}
				else {
					$("#rsvpwrap form").replaceWith("<h3>Doh! It seems the gremlins have done something. The server has vanished. Please come back later and try again. Worst-case snail mail wins.</h3>");
				}
			});
		});


		/* Hero height
		================================================== */
		var windowHeight = $(window).height();
		
		$('.hero').height( windowHeight );
		
		$(window).resize(function() {
			
			var windowHeight = $(window).height();
			$('.hero').height( windowHeight );
			
		});

		// Menu settings
		$('#menuToggle, .menu-close').on('click', function(){
			$('#menuToggle').toggleClass('active');
			$('body').toggleClass('body-push-toleft');
			$('#theMenu').toggleClass('menu-open');
		});
			
		/* Gallery
		================================================== */
		new Photostack( document.getElementById( 'photostack' ), {
			callback : function( item ) {
				//console.log(item)
			}
		} );	
			
			/* Gallery popup
		=================================================== */
		var photos = [
			"after-the-proposal",
			"anna-jason-wedding",
			"apple-picking-us",
			"audreys-party",
			"billies-wedding",
			"brewery",
			"charger-game-nj",
			"chophouse",
			"coti-1",
			"cp-1",
			"cp-2",
			"cp-3",
			"cp-4",
			"cp-5",
			"cp-6",
			"cruise-dinner",
			"dinner1",
			"du-27th-1",
			"du-27th-2",
			"du-grad",
			"engage-1",
			"engage-2",
			"engage-3",
			"engage-4",
			"engage-5",
			"engage-6",
			"erica-eli-wedding",
			"first-date-kinda-2",
			"first-date-kinda",
			"flat-iron",
			"hall-of-fame",
			"magic-mountain",
			"maine-1",
			"maine-2",
			"mi-tailgate",
			"navy-pier",
			"new-years-2008",
			"no-clue-1",
			"nyc-skating",
			"paris-1",
			"paris-2",
			"proposal",
			"rushmore",
			"sd-1",
			"seattle",
			"shake-shack",
			"sig-kappa-us",
			"snowboarding",
			"the-beginning",
			"tx-fishing",
			"vegas1",
			"vegas2",
			"yankee-game",
			"du-30th",
			"magic-mountain-2"
		];
		// Select the images to show in this gallery
		var selectedPhotos = []
		while(selectedPhotos.length != 12) {
			var random = Math.floor(Math.random() * photos.length);
			// Only use the photo once
			if(selectedPhotos.indexOf(photos[random]) === -1) {
				selectedPhotos.push(photos[random]);
			}
		}
		// Grab all of the anchors in the gallery and set their href attribute
		$('#photostack a').each(function(index) {
			$(this).attr('href', '/static/images/placeholders/600x500/' + selectedPhotos[index] + '.jpg');
		});
		// Grab all of the img tags and set their source and alt attributes
		$('#photostack img').each(function(index) {
			$(this).attr('src', '/static/images/placeholders/240x240/' + selectedPhotos[index] + '.jpg');
			$(this).attr('alt', selectedPhotos[index]);
		});
		$('.photostack').magnificPopup({
			delegate: 'a',
			type: 'image',
			tLoading: 'Loading image #%curr%...',
			mainClass: 'mfp-img-mobile',
			gallery: {
				enabled: true,
				navigateByImgClick: true,
				preload: [0,1] // Will preload 0 - before current, and 1 after the current image
			},
			image: {
				tError: '<a href="%url%">The image #%curr%</a> could not be loaded.',
				titleSrc: function(item) {
					return item.el.attr('title');
				}
			}
			/* zoom: {
				enabled: true,
				duration: 300 // don't foget to change the duration also in CSS
			} */
		});
		
		//Home Background slider
		jQuery.supersized({	
		slide_interval          :   3000,		// Length between transitions
		transition              :   1, 			// 0-None, 1-Fade, 2-Slide Top, 3-Slide Right, 4-Slide Bottom, 5-Slide Left, 6-Carousel Right, 7-Carousel Left
		transition_speed		:	700,		// Speed of transition				
		slide_links				:	'blank',	// Individual links for each slide (Options: false, 'num', 'name', 'blank')
		slides 					:  	[			// Slideshow Images
										{image : '/static/images/placeholders/slider-0.jpg'},
										{image : '/static/images/placeholders/slider-1.jpg'},  
										{image : '/static/images/placeholders/slider-2.jpg'},
										{image : '/static/images/placeholders/slider-3.jpg'},
										{image : '/static/images/placeholders/slider-4.jpg'},
										{image : '/static/images/placeholders/slider-5.jpg'}
									]
		});	
	
	
			
	});		


})(jQuery);