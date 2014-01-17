/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : struct.go

* Purpose :

* Creation Date :

* Last Modified : Wed 15 Jan 2014 06:06:27 AM UTC

* Created By kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package nagiosToJson

import ()

type Mainstat struct {
	Info          Info
	Programstatus Programstatus
	Hoststatus    map[string]*Hoststatus
}

type Info struct {
	Created           string `json:"created"`
	Last_update_check string `json:"last_update_check"`
	Last_version      string `json:"last_version"`
	New_version       string `json:"new_version"`
	Update_available  string `json:"update_available"`
	Version           string `json:"version"`
}

type Programstatus struct {
	Active_host_checks_enabled           string `json:"active_host_checks_enabled"`
	Active_ondemand_host_check_stats     string `json:"active_ondemand_host_check_stats"`
	Active_ondemand_service_check_stats  string `json:"active_ondemand_service_check_stats"`
	Active_scheduled_host_check_stats    string `json:"active_scheduled_host_check_stats"`
	Active_scheduled_service_check_stats string `json:"active_scheduled_service_check_stats"`
	Active_service_checks_enabled        string `json:"active_service_checks_enabled"`
	Cached_host_check_stats              string `json:"cached_host_check_stats"`
	Cached_service_check_stats           string `json:"cached_service_check_stats"`
	Check_host_freshness                 string `json:"check_host_freshness"`
	Check_service_freshness              string `json:"check_service_freshness"`
	Daemon_mode                          string `json:"daemon_mode"`
	Enable_event_handlers                string `json:"enable_event_handlers"`
	Enable_failure_prediction            string `json:"enable_failure_prediction"`
	Enable_flap_detection                string `json:"enable_flap_detection"`
	Enable_notifications                 string `json:"enable_notifications"`
	External_command_stats               string `json:"external_command_stats"`
	Global_host_event_handler            string `json:"global_host_event_handler"`
	Global_service_event_handler         string `json:"global_service_event_handler"`
	High_external_command_buffer_slots   string `json:"high_external_command_buffer_slots"`
	Last_command_check                   string `json:"last_command_check"`
	Last_log_rotation                    string `json:"last_log_rotation"`
	Modified_host_attributes             string `json:"modified_host_attributes"`
	Modified_service_attributes          string `json:"modified_service_attributes"`
	Nagios_pid                           string `json:"nagios_pid"`
	Next_comment_id                      string `json:"next_comment_id"`
	Next_downtime_id                     string `json:"next_downtime_id"`
	Next_event_id                        string `json:"next_event_id"`
	Next_notification_id                 string `json:"next_notification_id"`
	Next_problem_id                      string `json:"next_problem_id"`
	Obsess_over_hosts                    string `json:"obsess_over_hosts"`
	Obsess_over_services                 string `json:"obsess_over_services"`
	Parallel_host_check_stats            string `json:"parallel_host_check_stats"`
	Passive_host_check_stats             string `json:"passive_host_check_stats"`
	Passive_host_checks_enabled          string `json:"passive_host_checks_enabled"`
	Passive_service_check_stats          string `json:"passive_service_check_stats"`
	Passive_service_checks_enabled       string `json:"passive_service_checks_enabled"`
	Process_performance_data             string `json:"process_performance_data"`
	Program_start                        string `json:"program_start"`
	Serial_host_check_stats              string `json:"serial_host_check_stats"`
	Total_external_command_buffer_slots  string `json:"total_external_command_buffer_slots"`
	Used_external_command_buffer_slots   string `json:"used_external_command_buffer_slots"`
}

type Hoststatus struct {
	Acknowledgement_type          string                    `json:"acknowledgement_type"`
	Active_checks_enabled         string                    `json:"active_checks_enabled"`
	Check_command                 string                    `json:"check_command"`
	Check_execution_time          string                    `json:"check_execution_time"`
	Check_interval                string                    `json:"check_interval"`
	Check_latency                 string                    `json:"check_latency"`
	Check_options                 string                    `json:"check_options"`
	Check_period                  string                    `json:"check_period"`
	Check_type                    string                    `json:"check_type"`
	Current_attempt               string                    `json:"current_attempt"`
	Current_event_id              string                    `json:"current_event_id"`
	Current_notification_id       string                    `json:"current_notification_id"`
	Current_notification_number   string                    `json:"current_notification_number"`
	Current_problem_id            string                    `json:"current_problem_id"`
	Current_state                 string                    `json:"current_state"`
	Event_handler                 string                    `json:"event_handler"`
	Event_handler_enabled         string                    `json:"event_handler_enabled"`
	Failure_prediction_enabled    string                    `json:"failure_prediction_enabled"`
	Flap_detection_enabled        string                    `json:"flap_detection_enabled"`
	Has_been_checked              string                    `json:"has_been_checked"`
	Is_flapping                   string                    `json:"is_flapping"`
	Last_check                    string                    `json:"last_check"`
	Last_event_id                 string                    `json:"last_event_id"`
	Last_hard_state               string                    `json:"last_hard_state"`
	Last_hard_state_change        string                    `json:"last_hard_state_change"`
	Last_notification             string                    `json:"last_notification"`
	Last_problem_id               string                    `json:"last_problem_id"`
	Last_state_change             string                    `json:"last_state_change"`
	Last_time_down                string                    `json:"last_time_down"`
	Last_time_unreachable         string                    `json:"last_time_unreachable"`
	Last_time_up                  string                    `json:"last_time_up"`
	Last_update                   string                    `json:"last_update"`
	Long_plugin_output            string                    `json:"long_plugin_output"`
	Max_attempts                  string                    `json:"max_attempts"`
	Modified_attributes           string                    `json:"modified_attributes"`
	Next_check                    string                    `json:"next_check"`
	Next_notification             string                    `json:"next_notification"`
	No_more_notifications         string                    `json:"no_more_notifications"`
	Notification_period           string                    `json:"notification_period"`
	Notifications_enabled         string                    `json:"notifications_enabled"`
	Obsess_over_host              string                    `json:"obsess_over_host"`
	Passive_checks_enabled        string                    `json:"passive_checks_enabled"`
	Percent_state_change          string                    `json:"percent_state_change"`
	Performance_data              string                    `json:"performance_data"`
	Plugin_output                 string                    `json:"plugin_output"`
	Problem_has_been_acknowledged string                    `json:"problem_has_been_acknowledged"`
	Process_performance_data      string                    `json:"process_performance_data"`
	Retry_interval                string                    `json:"retry_interval"`
	Scheduled_downtime_depth      string                    `json:"scheduled_downtime_depth"`
	Should_be_scheduled           string                    `json:"should_be_scheduled"`
	State_type                    string                    `json:"state_type"`
	Servicestatus                 map[string]*Servicestatus `json:"servicestatus"`
}

type Servicestatus struct {
	Acknowledgement_type          string `json:"acknowledgement_type"`
	Active_checks_enabled         string `json:"active_checks_enabled"`
	Check_command                 string `json:"check_command"`
	Check_execution_time          string `json:"check_execution_time"`
	Check_interval                string `json:"check_interval"`
	Check_latency                 string `json:"check_latency"`
	Check_options                 string `json:"check_options"`
	Check_period                  string `json:"check_period"`
	Check_type                    string `json:"check_type"`
	Current_attempt               string `json:"current_attempt"`
	Current_event_id              string `json:"current_event_id"`
	Current_notification_id       string `json:"current_notification_id"`
	Current_notification_number   string `json:"current_notification_number"`
	Current_problem_id            string `json:"current_problem_id"`
	Current_state                 string `json:"current_state"`
	Event_handler                 string `json:"event_handler"`
	Event_handler_enabled         string `json:"event_handler_enabled"`
	Failure_prediction_enabled    string `json:"failure_prediction_enabled"`
	Flap_detection_enabled        string `json:"flap_detection_enabled"`
	Has_been_checked              string `json:"has_been_checked"`
	Is_flapping                   string `json:"is_flapping"`
	Last_check                    string `json:"last_check"`
	Last_event_id                 string `json:"last_event_id"`
	Last_hard_state               string `json:"last_hard_state"`
	Last_hard_state_change        string `json:"last_hard_state_change"`
	Last_notification             string `json:"last_notification"`
	Last_problem_id               string `json:"last_problem_id"`
	Last_state_change             string `json:"last_state_change"`
	Last_time_critical            string `json:"last_time_critical"`
	Last_time_ok                  string `json:"last_time_ok"`
	Last_time_unknown             string `json:"last_time_unknown"`
	Last_time_warning             string `json:"last_time_warning"`
	Last_update                   string `json:"last_update"`
	Long_plugin_output            string `json:"long_plugin_output"`
	Max_attempts                  string `json:"max_attempts"`
	Modified_attributes           string `json:"modified_attributes"`
	Next_check                    string `json:"next_check"`
	Next_notification             string `json:"next_notification"`
	No_more_notifications         string `json:"no_more_notifications"`
	Notification_period           string `json:"notification_period"`
	Notifications_enabled         string `json:"notifications_enabled"`
	Obsess_over_service           string `json:"obsess_over_service"`
	Passive_checks_enabled        string `json:"passive_checks_enabled"`
	Percent_state_change          string `json:"percent_state_change"`
	Performance_data              string `json:"performance_data"`
	Plugin_output                 string `json:"plugin_output"`
	Problem_has_been_acknowledged string `json:"problem_has_been_acknowledged"`
	Process_performance_data      string `json:"process_performance_data"`
	Retry_interval                string `json:"retry_interval"`
	Scheduled_downtime_depth      string `json:"scheduled_downtime_depth"`
	Service_description           string `json:"service_description"`
	Should_be_scheduled           string `json:"should_be_scheduled"`
	State_type                    string `json:"state_type"`
}
