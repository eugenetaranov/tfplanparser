# tfplanparser
Parses terraform binary plan and prints it out

## Usage
./tfplanparser -path="path to terraform plan"

```
# ./tfplanparser -path=tfplan
aws_ebs_volume.mongodb.2 map[id:0xc4201bb680 availability_zone:0xc4201bb800 arn:0xc4201bb8c0 tags.%:0xc4201bbb00 iops:0xc4201bbd40 size:0xc4201bb2c0 tags.backup:0xc4201bb500 type:0xc4201bbbc0 tags.env:0xc4201bb740 snapshot_id:0xc4201bb980 kms_key_id:0xc4201bb380 tags.role:0xc4201bbc80 encrypted:0xc4201bba40 tags.project:0xc4201bb440 tags.Name:0xc4201bb5c0]
aws_ebs_volume.mongodb.0 map[tags.role:0xc4201bbe80 tags.%:0xc4201c0300 iops:0xc4201c06c0 type:0xc4201c0780 encrypted:0xc4201c0840 size:0xc4201bbf40 arn:0xc4201c0000 tags.Name:0xc4201c03c0 tags.backup:0xc4201c0540 tags.env:0xc4201c0600 tags.project:0xc4201c0240 kms_key_id:0xc4201c0900 snapshot_id:0xc4201c00c0 availability_zone:0xc4201c0180 id:0xc4201c0480]
aws_ebs_volume.mongodb.1 map[tags.Name:0xc4201c0a40 size:0xc4201c0f80 tags.backup:0xc4201c0d40 iops:0xc4201c0e00 type:0xc4201c0bc0 kms_key_id:0xc4201c0ec0 tags.env:0xc4201c1040 encrypted:0xc4201c1280 snapshot_id:0xc4201c1400 tags.role:0xc4201c0b00 tags.%:0xc4201c0c80 availability_zone:0xc4201c1100 arn:0xc4201c11c0 tags.project:0xc4201c1340 id:0xc4201c14c0]
aws_volume_attachment.mongodb.2 map[id:0xc4201c1780 volume_id:0xc4201c1840 force_detach:0xc4201c1940 skip_destroy:0xc4201c1a00 device_name:0xc4201c1600 instance_id:0xc4201c16c0]
aws_volume_attachment.mongodb.0 map[force_detach:0xc4201baac0 skip_destroy:0xc4201bab80 id:0xc4201bac40 device_name:0xc4201ba840 instance_id:0xc4201ba900 volume_id:0xc4201ba9c0]
aws_volume_attachment.mongodb.1 map[device_name:0xc4201bad80 instance_id:0xc4201bae40 volume_id:0xc4201baf00 force_detach:0xc4201bb000 skip_destroy:0xc4201bb0c0 id:0xc4201bb180]
```
