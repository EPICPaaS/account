

$(document).ready(function() {
	if ($('#error').length == 0) {
		$('#error').remove();
	};
	if ($('#error').length > 0) {
		$('#error').delay(2000).animate({left: '-280px', opacity: 0}, function() {
			$('#error').remove();
		});
	};
});