(function($){
	
	"use strict";
	
    $(document).ready(function(){

    	if (!String.prototype.startsWith) {
		  Object.defineProperty(String.prototype, 'startsWith', {
		    enumerable: false,
		    configurable: false,
		    writable: false,
		    value: function(searchString, position) {
		      position = position || 0;
		      return this.lastIndexOf(searchString, position) === position;
		    }
		  });
		}

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

		var validate = false;
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

		var validateGuest = function(guestIndex) {
			var valid = true;
			var firstName = $.trim($("#gFirst" + guestIndex).val());
			var attending = $("#gAttending" + guestIndex).val();
			// If no name and the default 'In Spirit' attendence is selected
			// assume that there isn't a guest provided for the allotted slot.
			if (firstName === "" && attending === "-1") return valid;
			// Make sure first name provided
			if (!!firstName) {
				$("#error-gFirst" + guestIndex).hide();
				$("#gFirst" + guestIndex).val(firstName);
			}
			else {
				$("#error-gFirst" + guestIndex).show();
				valid = false;
			}
			// Make sure attendence has been specified
			if (attending !== "-1") {
				$("#error-attending" + guestIndex).hide();
			}
			else {
				$("#error-attending" + guestIndex).show();
				valid = false;
			}
			if (attending == "1") {
				// Check that dinner selection has been made for the guest
				if ($("#gFood" + guestIndex).val() !== "-1") {
					$("#error-food" + guestIndex).hide();
				}
				else {
					$("#error-food" + guestIndex).show();
					valid = false;
				}
			}
			return valid;
		};

		var validateRSVPReply = function() {
			var valid = true;
			$("#rsvpwrap li").each(function(i, elem) {
				var jElem = $(elem);
				var anchor = jElem.children("a");
				var guestBoxId = anchor.attr("href");
				var guestIndex = guestBoxId.charAt(guestBoxId.length - 1);
				valid = validateGuest(guestIndex);
				if (valid) {
					anchor.removeClass("guestError");
				}
				else {
					anchor.addClass("guestError");
				}
			});
			return valid;
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
					$("#rsvp-reply-form").replaceWith(data);
					$(".error").hide();
				}
				else {
					$("#rsvp-reply-form").replaceWith("<h3>Doh! It seems the gremlins have done something. The server has vanished. Please come back later and try again. Worst-case snail mail wins.</h3>");
				}
			}).fail(function(jqXHR, textStatus, errorThrown) {
				var errHeader = $("#rsvpwrap h3");
				if (errHeader.length) {
					errHeader.remove();
				}
				if (!!jqXHR.responseText) {
					$("#rsvp-reply-form").replaceWith(jqXHR.responseText);
					$(".error").hide();
				}
				else {
					$("#rsvp-reply-form").replaceWith("<h3>Doh! It seems the gremlins have done something. The server has vanished. Please come back later and try again. Worst-case snail mail wins.</h3>");
				}
			});
		};

		$("#rsvpwrap").on("change", "select", function(e) {
			var selectBox = $(e.target);
			var selectWidget = selectBox.attr("id");
			var selectValue = selectBox.val();
			var selectIndex = selectWidget.charAt(selectWidget.length - 1);
			if (selectWidget.startsWith("gAttending")) {
				if (selectValue === "1") {
					$("#foodBox" + selectIndex).slideDown();
					$("#error-attending" + selectIndex).hide();
				}
				else if (selectValue === "0") {
					$("#foodBox" + selectIndex).slideUp();
					$("#gFood" + selectIndex).val("-1");
					$("#error-attending" + selectIndex).hide();
				}
			}
			else if (selectWidget.startsWith("gFood")) {
				if (selectValue !== "-1") {
					$("#error-food" + selectIndex).hide();
				}
			}
			if (validate) validateRSVPReply();
		});

		$("#rsvpwrap").on("focus", ".text-input, select", function () {
	        $(this).css({border:"2px solid #de675f"});
	        $(this).css({background:"#fff"});
	    });

	    $("#rsvpwrap").on("blur", ".text-input, select", function () {
	        $(this).css({border:"2px solid #fff"});
	        $(this).css({background:"transparent"});
	    });

	    // Pagination controls
	    $("#rsvpwrap").on("click", "a", function(e) {
	    	e.preventDefault();
	    	var newActive = $(e.target);
	    	var newGuest = $("#" + newActive.attr("href"));
	    	var oldActive = $("#rsvpwrap li.active");
	    	var oldGuest = $("#" + oldActive.children("a").attr("href"));
	    	oldActive.removeClass("active");
	    	newActive.parent().addClass("active");
	    	oldGuest.slideUp().fadeOut();
	    	newGuest.slideDown().fadeIn();
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
					$("#rsvp-reply-form").replaceWith("<h3>Doh! It seems the gremlins have done something. The server has vanished. Please come back later and try again. Worst-case snail mail wins.</h3>");
				}
			}).fail(function(jqXHR, textStatus, errorThrown) {
				var errHeader = $("#rsvpwrap h3");
				if (errHeader.length) {
					errHeader.remove();
				}
				if (!!jqXHR.responseText) {
					$("#rsvp-reply-form").replaceWith(jqXHR.responseText);
					$(".error").hide();
				}
				else {
					$("#rsvp-reply-form").replaceWith("<h3>Doh! It seems the gremlins have done something. The server has vanished. Please come back later and try again. Worst-case snail mail wins.</h3>");
				}
			});
		});

		var getRSVPPostData = function(form) {
			var allData = form.serializeArray();
			var completedData = [];
			// Need to keep the hidden values and the email so grab them first
			for (var i = 0; i < allData.length; i++) {
				var lowerName = (allData[i].name).toLowerCase();
				if (lowerName === "invitation.id" || lowerName === "hmac.hash" ||
					lowerName === "hmac.stamp" || lowerName === "invitation.confirmaddr.address") {
					completedData.push(allData[i]);
				}
			}
			// Grab guest values for those guest that info was provided
			$(".guestInfo").each(function(i) {
				var guestName = $.trim($("#gFirst" + i).val());
				if (guestName !== "") {
					var guestValues = allData.filter(function(elem) {
						var lowerName = (elem.name).toLowerCase();
						return lowerName.startsWith("invitation.guests." + i);
					});
					completedData.push.apply(completedData, guestValues);
				}
			});
			return ($.param(completedData)).toLowerCase();
		};

		$("#rsvpwrap").on("submit", "#rsvp-reply", function(e) {
			e.preventDefault();
			var validForm = validateRSVPReply();
			validate = !validForm;
			if (!validForm) return;
			$("#reply-submit-btn").prop("disabled", true);
			var form = $(this);
			var postData = getRSVPPostData(form);
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
					$("#rsvp-reply-form").replaceWith("<h1>Thank you for taking the time to RSVP!</h1>");
					$("a[href='#rsvp']").trigger("click");
				}
				else {
					$("#rsvp-reply-form").replaceWith("<h3>Doh! It seems the gremlins have done something. Please come back later and try again. Worst-case snail mail wins.</h3>");
				}
			}).fail(function(jqXHR, textStatus, errorThrown) {
				var errHeader = $("#rsvpwrap h3");
				if (errHeader.length) {
					errHeader.remove();
				}
				if (!!jqXHR.responseText) {
					$("#rsvp-reply-form").replaceWith(jqXHR.responseText);
					$(".error").hide();
				}
				else {
					$("#rsvp-reply-form").replaceWith("<h3>Doh! It seems the gremlins have done something. The server has vanished. Please come back later and try again. Worst-case snail mail wins.</h3>");
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