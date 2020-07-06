// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/jeffsvajlenko/fortissimo/server/ent/song"
	"github.com/jeffsvajlenko/fortissimo/server/ent/tag"
)

// SongCreate is the builder for creating a Song entity.
type SongCreate struct {
	config
	mutation *SongMutation
	hooks    []Hook
}

// SetPath sets the path field.
func (sc *SongCreate) SetPath(s string) *SongCreate {
	sc.mutation.SetPath(s)
	return sc
}

// SetTitle sets the title field.
func (sc *SongCreate) SetTitle(s string) *SongCreate {
	sc.mutation.SetTitle(s)
	return sc
}

// SetNillableTitle sets the title field if the given value is not nil.
func (sc *SongCreate) SetNillableTitle(s *string) *SongCreate {
	if s != nil {
		sc.SetTitle(*s)
	}
	return sc
}

// SetTitleSort sets the title_sort field.
func (sc *SongCreate) SetTitleSort(s string) *SongCreate {
	sc.mutation.SetTitleSort(s)
	return sc
}

// SetNillableTitleSort sets the title_sort field if the given value is not nil.
func (sc *SongCreate) SetNillableTitleSort(s *string) *SongCreate {
	if s != nil {
		sc.SetTitleSort(*s)
	}
	return sc
}

// SetArtists sets the artists field.
func (sc *SongCreate) SetArtists(s []string) *SongCreate {
	sc.mutation.SetArtists(s)
	return sc
}

// SetFirstArtist sets the first_artist field.
func (sc *SongCreate) SetFirstArtist(s string) *SongCreate {
	sc.mutation.SetFirstArtist(s)
	return sc
}

// SetNillableFirstArtist sets the first_artist field if the given value is not nil.
func (sc *SongCreate) SetNillableFirstArtist(s *string) *SongCreate {
	if s != nil {
		sc.SetFirstArtist(*s)
	}
	return sc
}

// SetFirstArtistSort sets the first_artist_sort field.
func (sc *SongCreate) SetFirstArtistSort(s string) *SongCreate {
	sc.mutation.SetFirstArtistSort(s)
	return sc
}

// SetNillableFirstArtistSort sets the first_artist_sort field if the given value is not nil.
func (sc *SongCreate) SetNillableFirstArtistSort(s *string) *SongCreate {
	if s != nil {
		sc.SetFirstArtistSort(*s)
	}
	return sc
}

// SetFirstAlbumArtist sets the first_album_artist field.
func (sc *SongCreate) SetFirstAlbumArtist(s string) *SongCreate {
	sc.mutation.SetFirstAlbumArtist(s)
	return sc
}

// SetNillableFirstAlbumArtist sets the first_album_artist field if the given value is not nil.
func (sc *SongCreate) SetNillableFirstAlbumArtist(s *string) *SongCreate {
	if s != nil {
		sc.SetFirstAlbumArtist(*s)
	}
	return sc
}

// SetFirstAlbumArtistSort sets the first_album_artist_sort field.
func (sc *SongCreate) SetFirstAlbumArtistSort(s string) *SongCreate {
	sc.mutation.SetFirstAlbumArtistSort(s)
	return sc
}

// SetNillableFirstAlbumArtistSort sets the first_album_artist_sort field if the given value is not nil.
func (sc *SongCreate) SetNillableFirstAlbumArtistSort(s *string) *SongCreate {
	if s != nil {
		sc.SetFirstAlbumArtistSort(*s)
	}
	return sc
}

// SetAlbumArtist sets the album_artist field.
func (sc *SongCreate) SetAlbumArtist(s string) *SongCreate {
	sc.mutation.SetAlbumArtist(s)
	return sc
}

// SetNillableAlbumArtist sets the album_artist field if the given value is not nil.
func (sc *SongCreate) SetNillableAlbumArtist(s *string) *SongCreate {
	if s != nil {
		sc.SetAlbumArtist(*s)
	}
	return sc
}

// SetAlbum sets the album field.
func (sc *SongCreate) SetAlbum(s string) *SongCreate {
	sc.mutation.SetAlbum(s)
	return sc
}

// SetNillableAlbum sets the album field if the given value is not nil.
func (sc *SongCreate) SetNillableAlbum(s *string) *SongCreate {
	if s != nil {
		sc.SetAlbum(*s)
	}
	return sc
}

// SetPublisher sets the publisher field.
func (sc *SongCreate) SetPublisher(s string) *SongCreate {
	sc.mutation.SetPublisher(s)
	return sc
}

// SetNillablePublisher sets the publisher field if the given value is not nil.
func (sc *SongCreate) SetNillablePublisher(s *string) *SongCreate {
	if s != nil {
		sc.SetPublisher(*s)
	}
	return sc
}

// SetFirstComposer sets the first_composer field.
func (sc *SongCreate) SetFirstComposer(s string) *SongCreate {
	sc.mutation.SetFirstComposer(s)
	return sc
}

// SetNillableFirstComposer sets the first_composer field if the given value is not nil.
func (sc *SongCreate) SetNillableFirstComposer(s *string) *SongCreate {
	if s != nil {
		sc.SetFirstComposer(*s)
	}
	return sc
}

// SetComposers sets the composers field.
func (sc *SongCreate) SetComposers(s string) *SongCreate {
	sc.mutation.SetComposers(s)
	return sc
}

// SetNillableComposers sets the composers field if the given value is not nil.
func (sc *SongCreate) SetNillableComposers(s *string) *SongCreate {
	if s != nil {
		sc.SetComposers(*s)
	}
	return sc
}

// SetConductor sets the conductor field.
func (sc *SongCreate) SetConductor(s string) *SongCreate {
	sc.mutation.SetConductor(s)
	return sc
}

// SetNillableConductor sets the conductor field if the given value is not nil.
func (sc *SongCreate) SetNillableConductor(s *string) *SongCreate {
	if s != nil {
		sc.SetConductor(*s)
	}
	return sc
}

// SetGenre sets the genre field.
func (sc *SongCreate) SetGenre(s string) *SongCreate {
	sc.mutation.SetGenre(s)
	return sc
}

// SetNillableGenre sets the genre field if the given value is not nil.
func (sc *SongCreate) SetNillableGenre(s *string) *SongCreate {
	if s != nil {
		sc.SetGenre(*s)
	}
	return sc
}

// SetGrouping sets the grouping field.
func (sc *SongCreate) SetGrouping(s string) *SongCreate {
	sc.mutation.SetGrouping(s)
	return sc
}

// SetNillableGrouping sets the grouping field if the given value is not nil.
func (sc *SongCreate) SetNillableGrouping(s *string) *SongCreate {
	if s != nil {
		sc.SetGrouping(*s)
	}
	return sc
}

// SetYear sets the year field.
func (sc *SongCreate) SetYear(u uint) *SongCreate {
	sc.mutation.SetYear(u)
	return sc
}

// SetNillableYear sets the year field if the given value is not nil.
func (sc *SongCreate) SetNillableYear(u *uint) *SongCreate {
	if u != nil {
		sc.SetYear(*u)
	}
	return sc
}

// SetTrackNumber sets the track_number field.
func (sc *SongCreate) SetTrackNumber(u uint) *SongCreate {
	sc.mutation.SetTrackNumber(u)
	return sc
}

// SetNillableTrackNumber sets the track_number field if the given value is not nil.
func (sc *SongCreate) SetNillableTrackNumber(u *uint) *SongCreate {
	if u != nil {
		sc.SetTrackNumber(*u)
	}
	return sc
}

// SetOfTrackNumber sets the of_track_number field.
func (sc *SongCreate) SetOfTrackNumber(u uint) *SongCreate {
	sc.mutation.SetOfTrackNumber(u)
	return sc
}

// SetNillableOfTrackNumber sets the of_track_number field if the given value is not nil.
func (sc *SongCreate) SetNillableOfTrackNumber(u *uint) *SongCreate {
	if u != nil {
		sc.SetOfTrackNumber(*u)
	}
	return sc
}

// SetDiskNumber sets the disk_number field.
func (sc *SongCreate) SetDiskNumber(u uint) *SongCreate {
	sc.mutation.SetDiskNumber(u)
	return sc
}

// SetNillableDiskNumber sets the disk_number field if the given value is not nil.
func (sc *SongCreate) SetNillableDiskNumber(u *uint) *SongCreate {
	if u != nil {
		sc.SetDiskNumber(*u)
	}
	return sc
}

// SetOfDiskNumber sets the of_disk_number field.
func (sc *SongCreate) SetOfDiskNumber(u uint) *SongCreate {
	sc.mutation.SetOfDiskNumber(u)
	return sc
}

// SetNillableOfDiskNumber sets the of_disk_number field if the given value is not nil.
func (sc *SongCreate) SetNillableOfDiskNumber(u *uint) *SongCreate {
	if u != nil {
		sc.SetOfDiskNumber(*u)
	}
	return sc
}

// SetDuration sets the duration field.
func (sc *SongCreate) SetDuration(i int) *SongCreate {
	sc.mutation.SetDuration(i)
	return sc
}

// SetNillableDuration sets the duration field if the given value is not nil.
func (sc *SongCreate) SetNillableDuration(i *int) *SongCreate {
	if i != nil {
		sc.SetDuration(*i)
	}
	return sc
}

// SetPlayCount sets the play_count field.
func (sc *SongCreate) SetPlayCount(u uint) *SongCreate {
	sc.mutation.SetPlayCount(u)
	return sc
}

// SetNillablePlayCount sets the play_count field if the given value is not nil.
func (sc *SongCreate) SetNillablePlayCount(u *uint) *SongCreate {
	if u != nil {
		sc.SetPlayCount(*u)
	}
	return sc
}

// SetSkippedCount sets the skipped_count field.
func (sc *SongCreate) SetSkippedCount(u uint) *SongCreate {
	sc.mutation.SetSkippedCount(u)
	return sc
}

// SetNillableSkippedCount sets the skipped_count field if the given value is not nil.
func (sc *SongCreate) SetNillableSkippedCount(u *uint) *SongCreate {
	if u != nil {
		sc.SetSkippedCount(*u)
	}
	return sc
}

// SetComment sets the comment field.
func (sc *SongCreate) SetComment(s string) *SongCreate {
	sc.mutation.SetComment(s)
	return sc
}

// SetNillableComment sets the comment field if the given value is not nil.
func (sc *SongCreate) SetNillableComment(s *string) *SongCreate {
	if s != nil {
		sc.SetComment(*s)
	}
	return sc
}

// SetBeatsPerMinute sets the beats_per_minute field.
func (sc *SongCreate) SetBeatsPerMinute(u uint) *SongCreate {
	sc.mutation.SetBeatsPerMinute(u)
	return sc
}

// SetNillableBeatsPerMinute sets the beats_per_minute field if the given value is not nil.
func (sc *SongCreate) SetNillableBeatsPerMinute(u *uint) *SongCreate {
	if u != nil {
		sc.SetBeatsPerMinute(*u)
	}
	return sc
}

// SetCopyright sets the copyright field.
func (sc *SongCreate) SetCopyright(s string) *SongCreate {
	sc.mutation.SetCopyright(s)
	return sc
}

// SetNillableCopyright sets the copyright field if the given value is not nil.
func (sc *SongCreate) SetNillableCopyright(s *string) *SongCreate {
	if s != nil {
		sc.SetCopyright(*s)
	}
	return sc
}

// SetDateTagged sets the date_tagged field.
func (sc *SongCreate) SetDateTagged(t time.Time) *SongCreate {
	sc.mutation.SetDateTagged(t)
	return sc
}

// SetNillableDateTagged sets the date_tagged field if the given value is not nil.
func (sc *SongCreate) SetNillableDateTagged(t *time.Time) *SongCreate {
	if t != nil {
		sc.SetDateTagged(*t)
	}
	return sc
}

// SetDescription sets the description field.
func (sc *SongCreate) SetDescription(s string) *SongCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the description field if the given value is not nil.
func (sc *SongCreate) SetNillableDescription(s *string) *SongCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetFirstComposerSort sets the first_composer_sort field.
func (sc *SongCreate) SetFirstComposerSort(s string) *SongCreate {
	sc.mutation.SetFirstComposerSort(s)
	return sc
}

// SetNillableFirstComposerSort sets the first_composer_sort field if the given value is not nil.
func (sc *SongCreate) SetNillableFirstComposerSort(s *string) *SongCreate {
	if s != nil {
		sc.SetFirstComposerSort(*s)
	}
	return sc
}

// SetArtistsSort sets the artists_sort field.
func (sc *SongCreate) SetArtistsSort(s string) *SongCreate {
	sc.mutation.SetArtistsSort(s)
	return sc
}

// SetNillableArtistsSort sets the artists_sort field if the given value is not nil.
func (sc *SongCreate) SetNillableArtistsSort(s *string) *SongCreate {
	if s != nil {
		sc.SetArtistsSort(*s)
	}
	return sc
}

// SetLyrics sets the lyrics field.
func (sc *SongCreate) SetLyrics(s string) *SongCreate {
	sc.mutation.SetLyrics(s)
	return sc
}

// SetNillableLyrics sets the lyrics field if the given value is not nil.
func (sc *SongCreate) SetNillableLyrics(s *string) *SongCreate {
	if s != nil {
		sc.SetLyrics(*s)
	}
	return sc
}

// SetInitialKey sets the initial_key field.
func (sc *SongCreate) SetInitialKey(s string) *SongCreate {
	sc.mutation.SetInitialKey(s)
	return sc
}

// SetNillableInitialKey sets the initial_key field if the given value is not nil.
func (sc *SongCreate) SetNillableInitialKey(s *string) *SongCreate {
	if s != nil {
		sc.SetInitialKey(*s)
	}
	return sc
}

// SetIsrc sets the isrc field.
func (sc *SongCreate) SetIsrc(s string) *SongCreate {
	sc.mutation.SetIsrc(s)
	return sc
}

// SetNillableIsrc sets the isrc field if the given value is not nil.
func (sc *SongCreate) SetNillableIsrc(s *string) *SongCreate {
	if s != nil {
		sc.SetIsrc(*s)
	}
	return sc
}

// SetSubtitle sets the subtitle field.
func (sc *SongCreate) SetSubtitle(s string) *SongCreate {
	sc.mutation.SetSubtitle(s)
	return sc
}

// SetNillableSubtitle sets the subtitle field if the given value is not nil.
func (sc *SongCreate) SetNillableSubtitle(s *string) *SongCreate {
	if s != nil {
		sc.SetSubtitle(*s)
	}
	return sc
}

// SetMusicBrainzArtistID sets the music_brainz_artist_id field.
func (sc *SongCreate) SetMusicBrainzArtistID(s string) *SongCreate {
	sc.mutation.SetMusicBrainzArtistID(s)
	return sc
}

// SetNillableMusicBrainzArtistID sets the music_brainz_artist_id field if the given value is not nil.
func (sc *SongCreate) SetNillableMusicBrainzArtistID(s *string) *SongCreate {
	if s != nil {
		sc.SetMusicBrainzArtistID(*s)
	}
	return sc
}

// SetMusicBrainzDiscID sets the music_brainz_disc_id field.
func (sc *SongCreate) SetMusicBrainzDiscID(s string) *SongCreate {
	sc.mutation.SetMusicBrainzDiscID(s)
	return sc
}

// SetNillableMusicBrainzDiscID sets the music_brainz_disc_id field if the given value is not nil.
func (sc *SongCreate) SetNillableMusicBrainzDiscID(s *string) *SongCreate {
	if s != nil {
		sc.SetMusicBrainzDiscID(*s)
	}
	return sc
}

// SetMusicBrainzReleaseArtistID sets the music_brainz_release_artist_id field.
func (sc *SongCreate) SetMusicBrainzReleaseArtistID(s string) *SongCreate {
	sc.mutation.SetMusicBrainzReleaseArtistID(s)
	return sc
}

// SetNillableMusicBrainzReleaseArtistID sets the music_brainz_release_artist_id field if the given value is not nil.
func (sc *SongCreate) SetNillableMusicBrainzReleaseArtistID(s *string) *SongCreate {
	if s != nil {
		sc.SetMusicBrainzReleaseArtistID(*s)
	}
	return sc
}

// SetMusicBrainzReleaseCountry sets the music_brainz_release_country field.
func (sc *SongCreate) SetMusicBrainzReleaseCountry(s string) *SongCreate {
	sc.mutation.SetMusicBrainzReleaseCountry(s)
	return sc
}

// SetNillableMusicBrainzReleaseCountry sets the music_brainz_release_country field if the given value is not nil.
func (sc *SongCreate) SetNillableMusicBrainzReleaseCountry(s *string) *SongCreate {
	if s != nil {
		sc.SetMusicBrainzReleaseCountry(*s)
	}
	return sc
}

// SetMusicBrainzReleaseGroupID sets the music_brainz_release_group_id field.
func (sc *SongCreate) SetMusicBrainzReleaseGroupID(s string) *SongCreate {
	sc.mutation.SetMusicBrainzReleaseGroupID(s)
	return sc
}

// SetNillableMusicBrainzReleaseGroupID sets the music_brainz_release_group_id field if the given value is not nil.
func (sc *SongCreate) SetNillableMusicBrainzReleaseGroupID(s *string) *SongCreate {
	if s != nil {
		sc.SetMusicBrainzReleaseGroupID(*s)
	}
	return sc
}

// SetMusicBrainzReleaseID sets the music_brainz_release_id field.
func (sc *SongCreate) SetMusicBrainzReleaseID(s string) *SongCreate {
	sc.mutation.SetMusicBrainzReleaseID(s)
	return sc
}

// SetNillableMusicBrainzReleaseID sets the music_brainz_release_id field if the given value is not nil.
func (sc *SongCreate) SetNillableMusicBrainzReleaseID(s *string) *SongCreate {
	if s != nil {
		sc.SetMusicBrainzReleaseID(*s)
	}
	return sc
}

// SetMusicBrainzReleaseStatus sets the music_brainz_release_status field.
func (sc *SongCreate) SetMusicBrainzReleaseStatus(s string) *SongCreate {
	sc.mutation.SetMusicBrainzReleaseStatus(s)
	return sc
}

// SetNillableMusicBrainzReleaseStatus sets the music_brainz_release_status field if the given value is not nil.
func (sc *SongCreate) SetNillableMusicBrainzReleaseStatus(s *string) *SongCreate {
	if s != nil {
		sc.SetMusicBrainzReleaseStatus(*s)
	}
	return sc
}

// SetMusicBrainzReleaseType sets the music_brainz_release_type field.
func (sc *SongCreate) SetMusicBrainzReleaseType(s string) *SongCreate {
	sc.mutation.SetMusicBrainzReleaseType(s)
	return sc
}

// SetNillableMusicBrainzReleaseType sets the music_brainz_release_type field if the given value is not nil.
func (sc *SongCreate) SetNillableMusicBrainzReleaseType(s *string) *SongCreate {
	if s != nil {
		sc.SetMusicBrainzReleaseType(*s)
	}
	return sc
}

// SetMusicBrainzTrackID sets the music_brainz_track_id field.
func (sc *SongCreate) SetMusicBrainzTrackID(s string) *SongCreate {
	sc.mutation.SetMusicBrainzTrackID(s)
	return sc
}

// SetNillableMusicBrainzTrackID sets the music_brainz_track_id field if the given value is not nil.
func (sc *SongCreate) SetNillableMusicBrainzTrackID(s *string) *SongCreate {
	if s != nil {
		sc.SetMusicBrainzTrackID(*s)
	}
	return sc
}

// SetMusicIPID sets the music_ip_id field.
func (sc *SongCreate) SetMusicIPID(s string) *SongCreate {
	sc.mutation.SetMusicIPID(s)
	return sc
}

// SetNillableMusicIPID sets the music_ip_id field if the given value is not nil.
func (sc *SongCreate) SetNillableMusicIPID(s *string) *SongCreate {
	if s != nil {
		sc.SetMusicIPID(*s)
	}
	return sc
}

// SetRemixedBy sets the remixed_by field.
func (sc *SongCreate) SetRemixedBy(s string) *SongCreate {
	sc.mutation.SetRemixedBy(s)
	return sc
}

// SetNillableRemixedBy sets the remixed_by field if the given value is not nil.
func (sc *SongCreate) SetNillableRemixedBy(s *string) *SongCreate {
	if s != nil {
		sc.SetRemixedBy(*s)
	}
	return sc
}

// SetReplayGainAlbumGain sets the replay_gain_album_gain field.
func (sc *SongCreate) SetReplayGainAlbumGain(f float64) *SongCreate {
	sc.mutation.SetReplayGainAlbumGain(f)
	return sc
}

// SetNillableReplayGainAlbumGain sets the replay_gain_album_gain field if the given value is not nil.
func (sc *SongCreate) SetNillableReplayGainAlbumGain(f *float64) *SongCreate {
	if f != nil {
		sc.SetReplayGainAlbumGain(*f)
	}
	return sc
}

// SetReplayGainAlbumPeak sets the replay_gain_album_peak field.
func (sc *SongCreate) SetReplayGainAlbumPeak(f float64) *SongCreate {
	sc.mutation.SetReplayGainAlbumPeak(f)
	return sc
}

// SetNillableReplayGainAlbumPeak sets the replay_gain_album_peak field if the given value is not nil.
func (sc *SongCreate) SetNillableReplayGainAlbumPeak(f *float64) *SongCreate {
	if f != nil {
		sc.SetReplayGainAlbumPeak(*f)
	}
	return sc
}

// SetReplayGainTrackGain sets the replay_gain_track_gain field.
func (sc *SongCreate) SetReplayGainTrackGain(f float64) *SongCreate {
	sc.mutation.SetReplayGainTrackGain(f)
	return sc
}

// SetNillableReplayGainTrackGain sets the replay_gain_track_gain field if the given value is not nil.
func (sc *SongCreate) SetNillableReplayGainTrackGain(f *float64) *SongCreate {
	if f != nil {
		sc.SetReplayGainTrackGain(*f)
	}
	return sc
}

// SetReplayGainTrackPeak sets the replay_gain_track_peak field.
func (sc *SongCreate) SetReplayGainTrackPeak(f float64) *SongCreate {
	sc.mutation.SetReplayGainTrackPeak(f)
	return sc
}

// SetNillableReplayGainTrackPeak sets the replay_gain_track_peak field if the given value is not nil.
func (sc *SongCreate) SetNillableReplayGainTrackPeak(f *float64) *SongCreate {
	if f != nil {
		sc.SetReplayGainTrackPeak(*f)
	}
	return sc
}

// SetMimeType sets the mime_type field.
func (sc *SongCreate) SetMimeType(s string) *SongCreate {
	sc.mutation.SetMimeType(s)
	return sc
}

// SetNillableMimeType sets the mime_type field if the given value is not nil.
func (sc *SongCreate) SetNillableMimeType(s *string) *SongCreate {
	if s != nil {
		sc.SetMimeType(*s)
	}
	return sc
}

// AddTagIDs adds the tags edge to Tag by ids.
func (sc *SongCreate) AddTagIDs(ids ...int) *SongCreate {
	sc.mutation.AddTagIDs(ids...)
	return sc
}

// AddTags adds the tags edges to Tag.
func (sc *SongCreate) AddTags(t ...*Tag) *SongCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTagIDs(ids...)
}

// Mutation returns the SongMutation object of the builder.
func (sc *SongCreate) Mutation() *SongMutation {
	return sc.mutation
}

// Save creates the Song in the database.
func (sc *SongCreate) Save(ctx context.Context) (*Song, error) {
	if _, ok := sc.mutation.Path(); !ok {
		return nil, &ValidationError{Name: "path", err: errors.New("ent: missing required field \"path\"")}
	}
	if _, ok := sc.mutation.PlayCount(); !ok {
		v := song.DefaultPlayCount
		sc.mutation.SetPlayCount(v)
	}
	if _, ok := sc.mutation.SkippedCount(); !ok {
		v := song.DefaultSkippedCount
		sc.mutation.SetSkippedCount(v)
	}
	var (
		err  error
		node *Song
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SongMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SongCreate) SaveX(ctx context.Context) *Song {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SongCreate) sqlSave(ctx context.Context) (*Song, error) {
	var (
		s     = &Song{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: song.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: song.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldPath,
		})
		s.Path = value
	}
	if value, ok := sc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldTitle,
		})
		s.Title = value
	}
	if value, ok := sc.mutation.TitleSort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldTitleSort,
		})
		s.TitleSort = value
	}
	if value, ok := sc.mutation.Artists(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: song.FieldArtists,
		})
		s.Artists = value
	}
	if value, ok := sc.mutation.FirstArtist(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldFirstArtist,
		})
		s.FirstArtist = value
	}
	if value, ok := sc.mutation.FirstArtistSort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldFirstArtistSort,
		})
		s.FirstArtistSort = value
	}
	if value, ok := sc.mutation.FirstAlbumArtist(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldFirstAlbumArtist,
		})
		s.FirstAlbumArtist = value
	}
	if value, ok := sc.mutation.FirstAlbumArtistSort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldFirstAlbumArtistSort,
		})
		s.FirstAlbumArtistSort = value
	}
	if value, ok := sc.mutation.AlbumArtist(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldAlbumArtist,
		})
		s.AlbumArtist = value
	}
	if value, ok := sc.mutation.Album(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldAlbum,
		})
		s.Album = value
	}
	if value, ok := sc.mutation.Publisher(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldPublisher,
		})
		s.Publisher = value
	}
	if value, ok := sc.mutation.FirstComposer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldFirstComposer,
		})
		s.FirstComposer = value
	}
	if value, ok := sc.mutation.Composers(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldComposers,
		})
		s.Composers = value
	}
	if value, ok := sc.mutation.Conductor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldConductor,
		})
		s.Conductor = value
	}
	if value, ok := sc.mutation.Genre(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldGenre,
		})
		s.Genre = value
	}
	if value, ok := sc.mutation.Grouping(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldGrouping,
		})
		s.Grouping = value
	}
	if value, ok := sc.mutation.Year(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: song.FieldYear,
		})
		s.Year = value
	}
	if value, ok := sc.mutation.TrackNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: song.FieldTrackNumber,
		})
		s.TrackNumber = value
	}
	if value, ok := sc.mutation.OfTrackNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: song.FieldOfTrackNumber,
		})
		s.OfTrackNumber = value
	}
	if value, ok := sc.mutation.DiskNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: song.FieldDiskNumber,
		})
		s.DiskNumber = value
	}
	if value, ok := sc.mutation.OfDiskNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: song.FieldOfDiskNumber,
		})
		s.OfDiskNumber = value
	}
	if value, ok := sc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: song.FieldDuration,
		})
		s.Duration = value
	}
	if value, ok := sc.mutation.PlayCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: song.FieldPlayCount,
		})
		s.PlayCount = value
	}
	if value, ok := sc.mutation.SkippedCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: song.FieldSkippedCount,
		})
		s.SkippedCount = value
	}
	if value, ok := sc.mutation.Comment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldComment,
		})
		s.Comment = value
	}
	if value, ok := sc.mutation.BeatsPerMinute(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: song.FieldBeatsPerMinute,
		})
		s.BeatsPerMinute = value
	}
	if value, ok := sc.mutation.Copyright(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldCopyright,
		})
		s.Copyright = value
	}
	if value, ok := sc.mutation.DateTagged(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: song.FieldDateTagged,
		})
		s.DateTagged = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldDescription,
		})
		s.Description = value
	}
	if value, ok := sc.mutation.FirstComposerSort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldFirstComposerSort,
		})
		s.FirstComposerSort = value
	}
	if value, ok := sc.mutation.ArtistsSort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldArtistsSort,
		})
		s.ArtistsSort = value
	}
	if value, ok := sc.mutation.Lyrics(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldLyrics,
		})
		s.Lyrics = value
	}
	if value, ok := sc.mutation.InitialKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldInitialKey,
		})
		s.InitialKey = value
	}
	if value, ok := sc.mutation.Isrc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldIsrc,
		})
		s.Isrc = value
	}
	if value, ok := sc.mutation.Subtitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldSubtitle,
		})
		s.Subtitle = value
	}
	if value, ok := sc.mutation.MusicBrainzArtistID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldMusicBrainzArtistID,
		})
		s.MusicBrainzArtistID = value
	}
	if value, ok := sc.mutation.MusicBrainzDiscID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldMusicBrainzDiscID,
		})
		s.MusicBrainzDiscID = value
	}
	if value, ok := sc.mutation.MusicBrainzReleaseArtistID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldMusicBrainzReleaseArtistID,
		})
		s.MusicBrainzReleaseArtistID = value
	}
	if value, ok := sc.mutation.MusicBrainzReleaseCountry(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldMusicBrainzReleaseCountry,
		})
		s.MusicBrainzReleaseCountry = value
	}
	if value, ok := sc.mutation.MusicBrainzReleaseGroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldMusicBrainzReleaseGroupID,
		})
		s.MusicBrainzReleaseGroupID = value
	}
	if value, ok := sc.mutation.MusicBrainzReleaseID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldMusicBrainzReleaseID,
		})
		s.MusicBrainzReleaseID = value
	}
	if value, ok := sc.mutation.MusicBrainzReleaseStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldMusicBrainzReleaseStatus,
		})
		s.MusicBrainzReleaseStatus = value
	}
	if value, ok := sc.mutation.MusicBrainzReleaseType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldMusicBrainzReleaseType,
		})
		s.MusicBrainzReleaseType = value
	}
	if value, ok := sc.mutation.MusicBrainzTrackID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldMusicBrainzTrackID,
		})
		s.MusicBrainzTrackID = value
	}
	if value, ok := sc.mutation.MusicIPID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldMusicIPID,
		})
		s.MusicIPID = value
	}
	if value, ok := sc.mutation.RemixedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldRemixedBy,
		})
		s.RemixedBy = value
	}
	if value, ok := sc.mutation.ReplayGainAlbumGain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: song.FieldReplayGainAlbumGain,
		})
		s.ReplayGainAlbumGain = value
	}
	if value, ok := sc.mutation.ReplayGainAlbumPeak(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: song.FieldReplayGainAlbumPeak,
		})
		s.ReplayGainAlbumPeak = value
	}
	if value, ok := sc.mutation.ReplayGainTrackGain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: song.FieldReplayGainTrackGain,
		})
		s.ReplayGainTrackGain = value
	}
	if value, ok := sc.mutation.ReplayGainTrackPeak(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: song.FieldReplayGainTrackPeak,
		})
		s.ReplayGainTrackPeak = value
	}
	if value, ok := sc.mutation.MimeType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldMimeType,
		})
		s.MimeType = value
	}
	if nodes := sc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   song.TagsTable,
			Columns: song.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}
