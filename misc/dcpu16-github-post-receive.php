<?php
  /*
   * GitHub Post-Receive Hook target.
   * http://badsyntax.co/post/github-web-hook-post-receive
   * For project dcpu16
   */

# Turn off error reporting
error_reporting(0);

try {
	# Decode the payload json string
	$payload = json_decode($_REQUEST['payload']);
} catch(Exception $e) {
	exit;
}

# Pushed to master?
if ($payload->ref === 'refs/heads/master') {
	# Log the payload object
	@file_put_contents('logs/github.txt', print_r($payload, TRUE), FILE_APPEND);

	# Prep the URL - replace https protocol with git protocol to prevent 'update-server-info' errors
	$url = str_replace('https://', 'git://', $payload->repository->url);

	# Run the build script as a background process
	`./build.sh ${url} {$payload->repository->name} > /dev/null 2>&1 &`;
}
?>