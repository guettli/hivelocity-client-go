# TicketCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subject** | **string** |  | 
**Source** | Pointer to **float32** |  | [optional] 
**Priority** | Pointer to **float32** |  | [optional] 
**Queue** | **string** |  | 
**Assignment** | Pointer to **float32** |  | [optional] 

## Methods

### NewTicketCreate

`func NewTicketCreate(subject string, queue string, ) *TicketCreate`

NewTicketCreate instantiates a new TicketCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketCreateWithDefaults

`func NewTicketCreateWithDefaults() *TicketCreate`

NewTicketCreateWithDefaults instantiates a new TicketCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *TicketCreate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TicketCreate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TicketCreate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TicketCreate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetStatus

`func (o *TicketCreate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TicketCreate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TicketCreate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TicketCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *TicketCreate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TicketCreate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TicketCreate) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetSource

`func (o *TicketCreate) GetSource() float32`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TicketCreate) GetSourceOk() (*float32, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TicketCreate) SetSource(v float32)`

SetSource sets Source field to given value.

### HasSource

`func (o *TicketCreate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPriority

`func (o *TicketCreate) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TicketCreate) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TicketCreate) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TicketCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQueue

`func (o *TicketCreate) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *TicketCreate) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *TicketCreate) SetQueue(v string)`

SetQueue sets Queue field to given value.


### GetAssignment

`func (o *TicketCreate) GetAssignment() float32`

GetAssignment returns the Assignment field if non-nil, zero value otherwise.

### GetAssignmentOk

`func (o *TicketCreate) GetAssignmentOk() (*float32, bool)`

GetAssignmentOk returns a tuple with the Assignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignment

`func (o *TicketCreate) SetAssignment(v float32)`

SetAssignment sets Assignment field to given value.

### HasAssignment

`func (o *TicketCreate) HasAssignment() bool`

HasAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


