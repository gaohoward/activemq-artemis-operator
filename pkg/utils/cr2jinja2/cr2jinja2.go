package cr2jinja2

import (
	brokerv3alpha1 "github.com/artemiscloud/activemq-artemis-operator/pkg/apis/broker/v3alpha1"
	//k8syaml "k8s.io/apimachinery/pkg/util/yaml"
	"strings"
	"strconv"
	"fmt"
	"io/ioutil"
)

func checkString(prop *string) *string {
	if nil == prop {
		return nil
	}
	return prop
}

func checkBool(prop *bool) *string {
	if nil == prop {
		return nil
	}
	tmp := strconv.FormatBool(*prop)
	return &tmp
}

func checkInt32(prop *int32) *string {
	if nil == prop {
		return nil
	}
	tmp := fmt.Sprint(*prop)
	return &tmp
}

func checkInt64(prop *int64) *string {
	if nil == prop {
		return nil
	}
	tmp := strconv.FormatInt(*prop, 10)
	return &tmp
}

/* return a yaml string */
func MakeBrokerCfgOverrides(customeResource *brokerv3alpha1.ActiveMQArtemis, envVar *string, output *string) string {

	var addressSettings *[]brokerv3alpha1.AddressSettingType = &customeResource.Spec.AddressSettings.AddressSetting
	var sb strings.Builder
	
	processAddressSettings(&sb, addressSettings);
	
	fmt.Println("envvar: " + (*envVar))
	if envVar != nil && *envVar != "" {
		//how to do?
	}

	result := sb.String()

	fmt.Println("output " + *output)
	if output != nil && *output != "" {
		err := ioutil.WriteFile(*output, []byte(result), 0644)
		if err != nil {
			panic(err)
		}
	}
	return result

}

func processAddressSettings(sb *strings.Builder, addressSettings *[]brokerv3alpha1.AddressSettingType) {
	
	if addressSettings == nil || len(*addressSettings) == 0 {
		return
	}
	sb.WriteString("  address_settings:\n")
	for _, s := range *addressSettings {
		sb.WriteString("  - match: " + s.Match + "\n")
		if value := checkString(s.DeadLetterAddress); value != nil {
			sb.WriteString("    dead_letter_address: " + *value + "\n")			
		}
		if value := checkBool(s.AutoCreateDeadLetterResources); value != nil {
			sb.WriteString("    auto_create_dead_letter_resources: " + *value + "\n")			
		}
		if value := checkString(s.DeadLetterQueuePrefix); value != nil {
			sb.WriteString("    dead_letter_queue_prefix: " + *value + "\n")			
		}
		if value := checkString(s.DeadLetterQueueSuffix); value != nil {
			sb.WriteString("    dead_letter_queue_suffix: " + *value + "\n")	
		}
		if value := checkString(s.ExpiryAddress); value != nil {
			sb.WriteString("    expiry_address: " + *value + "\n")
		}
		if value := checkBool(s.AutoCreateExpiryResources); value != nil {
			sb.WriteString("    auto_create_expiry_resources: " + *value + "\n")
		}
		if value := checkString(s.ExpiryQueuePrefix); value != nil {
			sb.WriteString("    expiry_queue_prefix: " + *value + "\n")
		}
		if value := checkString(s.ExpiryQueueSuffix); value != nil {
			sb.WriteString("    expiry_queue_suffix: " + *value + "\n")
		}
		if value := checkInt32(s.ExpiryDelay); value != nil {
			sb.WriteString("    expiry_delay: " + *value + "\n")
		}
		if value := checkInt32(s.MinExpiryDelay); value != nil {
			sb.WriteString("    min_expiry_delay: " + *value + "\n")
		}
		if value := checkInt32(s.MaxExpiryDelay); value != nil {
			sb.WriteString("    max_expiry_delay: " + *value + "\n")
		}
		if value := checkInt32(s.RedeliveryDelay); value != nil {
			sb.WriteString("    redelivery_delay: " + *value + "\n")
		}
		if value := checkInt32(s.RedeliveryDelayMultiplier); value != nil {
			sb.WriteString("    redelivery_delay_multiplier: " + *value + "\n")
		}
		if value := checkInt32(s.RedeliveryCollisionAvoidanceFactor); value != nil {
			sb.WriteString("    redelivery_collision_avoidance_factor: " + *value + "\n")
		}
		if value := checkInt32(s.MaxRedeliveryDelay); value != nil {
			sb.WriteString("    max_redelivery_delay: " + *value + "\n")
		}
		if value := checkInt32(s.MaxDeliveryAttempts); value != nil {
			sb.WriteString("    max_delivery_attempts: " + *value + "\n")
		}
		if value := checkInt64(s.MaxSizeBytes); value != nil {
			sb.WriteString("    max_size_bytes: " + *value + "\n")
		}
		if value := checkInt32(s.MaxSizeBytesRejectThreshold); value != nil {
			sb.WriteString("    max_size_bytes_reject_threshold: " + *value + "\n")
		}
		if value := checkInt64(s.PageSizeBytes); value != nil {
			sb.WriteString("    page_size_bytes: " + *value + "\n")
		}
		if value := checkInt32(s.PageMaxCacheSize); value != nil {
			sb.WriteString("    page_max_cache_size: " + *value + "\n")
		}
		if value := checkString(s.AddressFullPolicy); value != nil {
			sb.WriteString("    address_full_policy: " + *value + "\n")
		}
		if value := checkInt32(s.MessageCounterHistoryDayLimit); value != nil {
			sb.WriteString("    message_counter_history_day_limit: " + *value + "\n")
		}
		if value := checkBool(s.LastValueQueue); value != nil {
			sb.WriteString("    last_value_queue: " + *value + "\n")
		}
		if value := checkBool(s.DefaultLastValueQueue); value != nil {
			sb.WriteString("    default_last_value_queue: " + *value + "\n")
		}
		if value := checkString(s.DefaultLastValueKey); value != nil {
			sb.WriteString("    default_last_value_key: " + *value + "\n")
		}
		if value := checkBool(s.DefaultNonDestructive); value != nil {
			sb.WriteString("    default_non_destructive: " + *value + "\n")
		}
		if value := checkBool(s.DefaultExclusiveQueue); value != nil {
			sb.WriteString("    default_exclusive_queue: " + *value + "\n")
		}
		if value := checkBool(s.DefaultGroupRebalance); value != nil {
			sb.WriteString("    default_group_rebalance: " + *value + "\n")
		}
		if value := checkInt32(s.DefaultGroupBuckets); value != nil {
			sb.WriteString("    default_group_buckets: " + *value + "\n")
		}
		if value := checkString(s.DefaultGroupFirstKey); value != nil {
			sb.WriteString("    default_group_first_key: " + *value + "\n")
		}
		if value := checkInt32(s.DefaultConsumersBeforeDispatch); value != nil {
			sb.WriteString("    default_consumers_before_dispatch: " + *value + "\n")
		}
		if value := checkInt32(s.DefaultDelayBeforeDispatch); value != nil {
			sb.WriteString("    default_delay_before_dispatch: " + *value + "\n")
		}
		if value := checkInt32(s.RedistributionDelay); value != nil {
			sb.WriteString("    redistribution_delay: " + *value + "\n")
		}
		if value := checkBool(s.SendToDlaOnNoRoute); value != nil {
			sb.WriteString("    send_to_dla_on_no_route: " + *value + "\n")
		}
		if value := checkInt32(s.SlowConsumerThreshold); value != nil {
			sb.WriteString("    slow_consumer_threshold: " + *value + "\n")
		}
		if value := checkString(s.SlowConsumerPolicy); value != nil {
			sb.WriteString("    slow_consumer_policy: " + *value + "\n")
		}
		if value := checkInt32(s.SlowConsumerCheckPeriod); value != nil {
			sb.WriteString("    slow_consumer_check_period: " + *value + "\n")
		}
		if value := checkBool(s.AutoCreateJmsQueues); value != nil {
			sb.WriteString("    auto_create_jms_queues: " + *value + "\n")
		}
		if value := checkBool(s.AutoDeleteJmsQueues); value != nil {
			sb.WriteString("    auto_delete_jms_queues: " + *value + "\n")
		}
		if value := checkBool(s.AutoCreateJmsTopics); value != nil {
			sb.WriteString("    auto_create_jms_topics: " + *value + "\n")
		}
		if value := checkBool(s.AutoDeleteJmsTopics); value != nil {
			sb.WriteString("    auto_delete_jms_topics: " + *value + "\n")
		}
		if value := checkBool(s.AutoCreateQueues); value != nil {
			sb.WriteString("    auto_create_queues: " + *value + "\n")
		}
		if value := checkBool(s.AutoDeleteQueues); value != nil {
			sb.WriteString("    auto_delete_queues: " + *value + "\n")
		}
		if value := checkBool(s.AutoDeleteCreatedQueues); value != nil {
			sb.WriteString("    auto_delete_created_queues: " + *value + "\n")
		}
		if value := checkInt32(s.AutoDeleteQueuesDelay); value != nil {
			sb.WriteString("    auto_delete_queues_delay: " + *value + "\n")
		}
		if value := checkInt32(s.AutoDeleteQueuesMessageCount); value != nil {
			sb.WriteString("    auto_delete_queues_message_count: " + *value + "\n")
		}
		if value := checkString(s.ConfigDeleteQueues); value != nil {
			sb.WriteString("    config_delete_queues: " + *value + "\n")
		}
		if value := checkBool(s.AutoCreateAddresses); value != nil {
			sb.WriteString("    auto_create_addresses: " + *value + "\n")
		}
		if value := checkBool(s.AutoDeleteAddresses); value != nil {
			sb.WriteString("    auto_delete_addresses: " + *value + "\n")
		}
		if value := checkInt32(s.AutoDeleteAddressesDelay); value != nil {
			sb.WriteString("    auto_delete_addresses_delay: " + *value + "\n")
		}
		if value := checkString(s.ConfigDeleteAddresses); value != nil {
			sb.WriteString("    config_delete_addresses: " + *value + "\n")
		}
		if value := checkInt32(s.ManagementBrowsePageSize); value != nil {
			sb.WriteString("    management_browse_page_size: " + *value + "\n")
		}
		if value := checkBool(s.DefaultPurgeOnNoConsumers); value != nil {
			sb.WriteString("    default_purge_on_no_consumers: " + *value + "\n")
		}
		if value := checkInt32(s.DefaultMaxConsumers); value != nil {
			sb.WriteString("    default_max_consumers: " + *value + "\n")
		}
		if value := checkString(s.DefaultQueueRoutingType); value != nil {
			sb.WriteString("    default_queue_routing_type: " + *value + "\n")
		}
		if value := checkString(s.DefaultAddressRoutingType); value != nil {
			sb.WriteString("    default_address_routing_type: " + *value + "\n")
		}
		if value := checkInt32(s.DefaultConsumerWindowSize); value != nil {
			sb.WriteString("    default_consumer_window_size: " + *value + "\n")
		}
		if value := checkInt32(s.DefaultRingSize); value != nil {
			sb.WriteString("    default_ring_size: " + *value + "\n")
		}
		if value := checkInt32(s.RetroactiveMessageCount); value != nil {
			sb.WriteString("    retroactive_message_count: " + *value + "\n")
		}
		if value := checkBool(s.EnableMetrics); value != nil {
			sb.WriteString("    enable_metrics: " + *value + "\n")
		}
	}

}