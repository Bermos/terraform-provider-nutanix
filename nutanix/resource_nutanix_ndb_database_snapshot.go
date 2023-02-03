package nutanix

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-nutanix/client/era"
	"github.com/terraform-providers/terraform-provider-nutanix/utils"
)

func resourceNutanixNDBDatabaseSnapshot() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNutanixNDBDatabaseSnapshotCreate,
		ReadContext:   resourceNutanixNDBDatabaseSnapshotRead,
		UpdateContext: resourceNutanixNDBDatabaseSnapshotUpdate,
		DeleteContext: resourceNutanixNDBDatabaseSnapshotDelete,
		Schema: map[string]*schema.Schema{
			"time_machine_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"time_machine_name"},
			},
			"time_machine_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"time_machine_id"},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"remove_schedule_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"expiry_date_timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Asia/Calcutta",
			},
			"replicate_to_clusters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			// computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"properties": dataSourceEraDatabaseProperties(),
			"tags":       dataSourceEraDBInstanceTags(),
			"snapshot_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nx_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"protection_domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_snapshot_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_info_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"applicable_types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"snapshot_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secure_info": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"info": dataSourceEraDatabaseInfo(),
						"deregister_info": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"from_timestamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"to_timestamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"replication_retry_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"last_replication_retyr_source_snapshot_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"async": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"stand_by": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"curation_retry_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"operations_using_snapshot": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"software_snapshot_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"software_database_snapshot": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dbserver_storage_metadata_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"santised": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"santised_from_snapshot_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"processed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"database_snapshot": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"from_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"to_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dbserver_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dbserver_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dbserver_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replicated_snapshots": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"software_snapshot": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"santised_snapshots": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snapshot_family": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snapshot_timestamp_date": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lcm_config": dataSourceEraLCMConfig(),
			"parent_snapshot": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"snapshot_size": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func resourceNutanixNDBDatabaseSnapshotCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*Client).Era

	req := &era.DatabaseSnapshotRequest{}
	snapshotName := ""
	tmsID, tok := d.GetOk("time_machine_id")
	tmsName, tnOk := d.GetOk("time_machine_name")

	if !tok && !tnOk {
		return diag.Errorf("Atleast one of time_machine_id or time_machine_name is required to perform clone")
	}

	if len(tmsName.(string)) > 0 {
		// call time machine API with value-type name
		res, err := conn.Service.GetTimeMachine(ctx, tmsID.(string), tmsName.(string))
		if err != nil {
			return diag.FromErr(err)
		}

		tmsID = *res.ID
	}

	if name, ok := d.GetOk("name"); ok {
		req.Name = utils.StringPtr(name.(string))
		snapshotName = utils.StringValue(req.Name)
	}

	if rm, ok := d.GetOk("remove_schedule_in_days"); ok {
		lcmConfig := &era.LCMConfigSnapshot{}
		snapshotLCM := &era.SnapshotLCMConfig{}
		expDetails := &era.DBExpiryDetails{}

		expDetails.ExpireInDays = utils.IntPtr(rm.(int))

		if tmzone, pk := d.GetOk("expiry_date_timezone"); pk {
			expDetails.ExpiryDateTimezone = utils.StringPtr(tmzone.(string))
		}

		snapshotLCM.ExpiryDetails = expDetails
		lcmConfig.SnapshotLCMConfig = snapshotLCM
		req.LcmConfig = lcmConfig
	}

	if rep, ok := d.GetOk("replicate_to_clusters"); ok && len(rep.([]interface{})) > 0 {
		repList := rep.([]interface{})

		for _, v := range repList {
			req.ReplicateToClusters = append(req.ReplicateToClusters, utils.StringPtr(v.(string)))
		}
	}

	// call the snapshot API

	resp, err := conn.Service.DatabaseSnapshot(ctx, tmsID.(string), req)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId(resp.Entityid)

	// Get Operation ID from response of snapshot and poll for the operation to get completed.
	opID := resp.Operationid
	if opID == "" {
		return diag.Errorf("error: operation ID is an empty string")
	}
	opReq := era.GetOperationRequest{
		OperationID: opID,
	}

	log.Printf("polling for operation with id: %s\n", opID)

	// Poll for operation here - Operation GET Call
	stateConf := &resource.StateChangeConf{
		Pending: []string{"PENDING"},
		Target:  []string{"COMPLETED", "FAILED"},
		Refresh: eraRefresh(ctx, conn, opReq),
		Timeout: d.Timeout(schema.TimeoutCreate),
		Delay:   eraDelay,
	}

	if _, errWaitTask := stateConf.WaitForStateContext(ctx); errWaitTask != nil {
		return diag.Errorf("error waiting for snapshot	 (%s) to create: %s", resp.Entityid, errWaitTask)
	}

	// Get all the Snapshots based on tms

	uniqueID := ""
	timeStamp := 0
	tmsResp, ter := conn.Service.ListSnapshots(ctx, resp.Entityid)
	if ter != nil {
		return diag.FromErr(ter)
	}
	for _, val := range *tmsResp {
		if snapshotName == utils.StringValue(val.Name) {
			if (int(*val.SnapshotTimeStampDate)) > timeStamp {
				uniqueID = utils.StringValue(val.ID)
				timeStamp = int(utils.Int64Value(val.SnapshotTimeStampDate))
			}
		}
	}
	d.SetId(uniqueID)
	log.Printf("NDB database snapshot with %s id created successfully", d.Id())
	return resourceNutanixNDBDatabaseSnapshotRead(ctx, d, meta)
}

func resourceNutanixNDBDatabaseSnapshotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*Client).Era

	// setting the default values for Get snapshot
	filterParams := &era.FilterParams{}
	filterParams.LoadReplicatedChildSnapshots = "false"
	filterParams.TimeZone = "UTC"

	resp, err := conn.Service.GetSnapshot(ctx, d.Id(), filterParams)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("owner_id", resp.OwnerID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", resp.Description); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("date_created", resp.DateCreated); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("date_modified", resp.DateModified); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("properties", flattenDBInstanceProperties(resp.Properties)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tags", flattenDBTags(resp.Tags)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("snapshot_uuid", resp.SnapshotUUID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("nx_cluster_id", resp.NxClusterID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("protection_domain_id", resp.ProtectionDomainID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("parent_snapshot_id", resp.ParentSnapshotID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("time_machine_id", resp.TimeMachineID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("database_node_id", resp.DatabaseNodeID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("app_info_version", resp.AppInfoVersion); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("status", resp.Status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("type", resp.Type); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("applicable_types", resp.ApplicableTypes); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("snapshot_timestamp", resp.SnapshotTimeStamp); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("software_snapshot_id", resp.SoftwareSnapshotID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("software_database_snapshot", resp.SoftwareDatabaseSnapshot); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("dbserver_storage_metadata_version", resp.DBServerStorageMetadataVersion); err != nil {
		return diag.FromErr(err)
	}

	// if err := d.Set("santised", resp.Sanitized); err != nil {
	// 	return diag.FromErr(err)
	// }

	if err := d.Set("santised_from_snapshot_id", resp.SanitisedFromSnapshotID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("timezone", resp.TimeZone); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("processed", resp.Processed); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("database_snapshot", resp.DatabaseSnapshot); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("from_timestamp", resp.FromTimeStamp); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("to_timestamp", resp.ToTimeStamp); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("dbserver_id", resp.DbserverID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("dbserver_name", resp.DbserverName); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("dbserver_ip", resp.DbserverIP); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("replicated_snapshots", resp.ReplicatedSnapshots); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("software_snapshot", resp.SoftwareSnapshot); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("santised_snapshots", resp.SanitisedSnapshots); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("snapshot_family", resp.SnapshotFamily); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("snapshot_timestamp_date", resp.SnapshotTimeStampDate); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("parent_snapshot", resp.ParentSnapshot); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("snapshot_size", resp.SnapshotSize); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("lcm_config", flattenDBLcmConfig(resp.LcmConfig)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("metadata", flattenClonedMetadata(resp.Metadata)); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceNutanixNDBDatabaseSnapshotUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*Client).Era

	updateReq := &era.UpdateSnapshotRequest{}

	if d.HasChange("name") {
		updateReq.Name = utils.StringPtr(d.Get("name").(string))
	}

	// reset the name is by default value provided
	updateReq.ResetName = true

	// API to update database snapshot

	resp, err := conn.Service.UpdateSnapshot(ctx, d.Id(), updateReq)
	if err != nil {
		return diag.FromErr(err)
	}

	if resp != nil {
		if err = d.Set("name", resp.Name); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("NDB database snapshot with %s id updated successfully", d.Id())
	return resourceNutanixNDBDatabaseSnapshotRead(ctx, d, meta)
}

func resourceNutanixNDBDatabaseSnapshotDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*Client).Era

	resp, err := conn.Service.DeleteSnapshot(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	opID := resp.Operationid

	opReq := era.GetOperationRequest{
		OperationID: opID,
	}

	log.Printf("polling for operation with id: %s\n", opID)

	// Poll for operation here - Operation GET Call
	stateConf := &resource.StateChangeConf{
		Pending: []string{"PENDING"},
		Target:  []string{"COMPLETED", "FAILED"},
		Refresh: eraRefresh(ctx, conn, opReq),
		Timeout: d.Timeout(schema.TimeoutCreate),
		Delay:   eraDelay,
	}

	if _, errWaitTask := stateConf.WaitForStateContext(ctx); errWaitTask != nil {
		return diag.Errorf("error waiting for snapshot	 (%s) to delete: %s", resp.Entityid, errWaitTask)
	}

	d.SetId("")
	return nil
}
