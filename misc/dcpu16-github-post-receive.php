<?php
  /*
   * GitHub Post-Receive Hook target
   * For project dcpu16
   */
# Turn off error reporting
error_reporting(0);

file_put_contents('/tmp/hook.log', posix_geteuid()."\n", FILE_APPEND);
# Check the incoming request IP
$github_ips = array('207.97.227.253', '50.57.128.197', '108.171.174.178');
if(in_array($_SERVER['REMOTE_ADDR'],$github_ips)) {
	# Run the build script as a background process
	#`./dcpu16.sh 'git://github.com/radekh/dcpu16.git' 'dcpu16' > /dev/null 2>&1 &`;
	shell_exec('/home/radek/bin/github-dcpu16-post-receive.sh >/tmp/post-receive.log 2>&1 &');
}
?>
